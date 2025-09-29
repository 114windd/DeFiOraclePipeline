package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/api"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/cache"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/fetcher"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/metrics"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/normalizer"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/publisher"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/storage"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/utils"
)

func main() {
	// Load configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize components
	cache, err := cache.NewCache(config.RedisURL)
	if err != nil {
		log.Fatalf("Failed to initialize cache: %v", err)
	}
	defer cache.Close()

	storage, err := storage.NewStorage(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer storage.Close()

	publisher, err := publisher.NewPublisher(config.NATSURL, config.NATSSubject)
	if err != nil {
		log.Fatalf("Failed to initialize publisher: %v", err)
	}
	defer publisher.Close()

	metrics := metrics.NewMetrics()
	normalizer := normalizer.NewDefaultNormalizer()
	fetcher := fetcher.NewFetcher(config.CoinGeckoURL, config.FetchTimeout)

	// Initialize API
	api := api.NewAPI(cache, storage, metrics)

	// Start the price fetcher service
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start price fetching goroutine
	go startPriceFetcher(ctx, fetcher, normalizer, cache, storage, publisher, metrics, config)

	// Start HTTP server
	go func() {
		log.Printf("Starting HTTP server on %s", config.GetServerAddr())
		if err := api.Run(config.GetServerAddr()); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	cancel()

	// Give services time to shut down gracefully
	time.Sleep(2 * time.Second)
	log.Println("Server stopped")
}

// startPriceFetcher runs the price fetching service
func startPriceFetcher(
	ctx context.Context,
	fetcher *fetcher.Fetcher,
	normalizer *normalizer.Normalizer,
	cache *cache.Cache,
	storage *storage.Storage,
	publisher *publisher.Publisher,
	metrics *metrics.Metrics,
	config *utils.Config,
) {
	ticker := time.NewTicker(config.FetchInterval)
	defer ticker.Stop()

	log.Printf("Starting price fetcher with interval %v", config.FetchInterval)

	for {
		select {
		case <-ctx.Done():
			log.Println("Price fetcher stopped")
			return
		case <-ticker.C:
			fetchAndProcessPrice(fetcher, normalizer, cache, storage, publisher, metrics, config)
		}
	}
}

// fetchAndProcessPrice fetches a price and processes it through the pipeline
func fetchAndProcessPrice(
	fetcher *fetcher.Fetcher,
	normalizer *normalizer.Normalizer,
	cache *cache.Cache,
	storage *storage.Storage,
	publisher *publisher.Publisher,
	metrics *metrics.Metrics,
	config *utils.Config,
) {
	start := time.Now()

	// Fetch price
	price, err := fetcher.FetchETHUSDPrice()
	if err != nil {
		metrics.RecordFetchError("coingecko", "fetch_failed")
		log.Printf("Failed to fetch price: %v", err)
		return
	}

	// Validate price
	if err := normalizer.ValidatePrice(price); err != nil {
		metrics.RecordFetchError("coingecko", "validation_failed")
		log.Printf("Price validation failed: %v", err)
		return
	}

	// Normalize price
	normalizedPrice := normalizer.NormalizePrice(price)
	timestamp := time.Now()

	// Record success metrics
	metrics.RecordFetchSuccess("coingecko")
	metrics.RecordFetchLatency(time.Since(start), "coingecko", "success")

	// Get last price for comparison
	lastPriceData, err := cache.GetCachedPrice()
	var lastPrice float64
	if err == nil {
		lastPrice = lastPriceData.Price
	}

	// Cache the price
	if err := cache.CachePrice(normalizedPrice, timestamp, "coingecko"); err != nil {
		metrics.RecordCacheError("redis", "set")
		log.Printf("Failed to cache price: %v", err)
	} else {
		metrics.RecordCacheHit("redis")
	}

	// Store in database
	if err := storage.SavePrice(normalizedPrice, timestamp, "coingecko"); err != nil {
		metrics.RecordDBError("insert", "price_records", "save_failed")
		log.Printf("Failed to save price to database: %v", err)
	} else {
		metrics.RecordDBOperation("insert", "price_records")
	}

	// Publish to NATS with filtering
	if err := publisher.PublishPriceWithFilter(normalizedPrice, timestamp, "coingecko", lastPrice, config.PriceChangeThreshold); err != nil {
		metrics.RecordNATSError("publish")
		log.Printf("Failed to publish price: %v", err)
	} else {
		metrics.RecordNATSPublished(config.NATSSubject)
		metrics.RecordPriceUpdate("coingecko", "published")
	}

	// Update price age metric
	metrics.RecordPriceAge(0, "coingecko") // Just updated, so age is 0

	log.Printf("Successfully processed price: $%.2f", normalizedPrice)
}

