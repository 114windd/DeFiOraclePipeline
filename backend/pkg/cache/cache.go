package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// PriceData represents cached price information
type PriceData struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source,omitempty"`
}

// Cache handles Redis operations for price caching
type Cache struct {
	client *redis.Client
	ctx    context.Context
}

// NewCache creates a new cache instance
func NewCache(redisURL string) (*Cache, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	// Test connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Cache{
		client: client,
		ctx:    ctx,
	}, nil
}

// NewCacheWithClient creates a cache instance with an existing Redis client
func NewCacheWithClient(client *redis.Client) *Cache {
	return &Cache{
		client: client,
		ctx:    context.Background(),
	}
}

// CachePrice stores the latest ETH/USD price in Redis
func (c *Cache) CachePrice(price float64, timestamp time.Time, source string) error {
	priceData := PriceData{
		Price:     price,
		Timestamp: timestamp,
		Source:    source,
	}

	data, err := json.Marshal(priceData)
	if err != nil {
		return fmt.Errorf("failed to marshal price data: %w", err)
	}

	key := "eth_usd_price"
	err = c.client.Set(c.ctx, key, data, 0).Err() // No expiration for latest price
	if err != nil {
		return fmt.Errorf("failed to cache price: %w", err)
	}

	return nil
}

// GetCachedPrice retrieves the most recent cached ETH/USD price
func (c *Cache) GetCachedPrice() (*PriceData, error) {
	key := "eth_usd_price"
	data, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no cached price found")
		}
		return nil, fmt.Errorf("failed to get cached price: %w", err)
	}

	var priceData PriceData
	if err := json.Unmarshal([]byte(data), &priceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached price: %w", err)
	}

	return &priceData, nil
}

// CachePriceHistory stores historical price data
func (c *Cache) CachePriceHistory(price float64, timestamp time.Time, source string) error {
	priceData := PriceData{
		Price:     price,
		Timestamp: timestamp,
		Source:    source,
	}

	data, err := json.Marshal(priceData)
	if err != nil {
		return fmt.Errorf("failed to marshal price data: %w", err)
	}

	// Use timestamp as part of the key for historical data
	key := fmt.Sprintf("eth_usd_price_history:%d", timestamp.Unix())

	// Store with 24 hour expiration for historical data
	err = c.client.Set(c.ctx, key, data, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to cache price history: %w", err)
	}

	return nil
}

// GetPriceHistory retrieves recent price records from cache
func (c *Cache) GetPriceHistory(limit int) ([]PriceData, error) {
	pattern := "eth_usd_price_history:*"
	keys, err := c.client.Keys(c.ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get price history keys: %w", err)
	}

	// Sort keys by timestamp (newest first) and limit results
	if len(keys) > limit {
		keys = keys[:limit]
	}

	var prices []PriceData
	for _, key := range keys {
		data, err := c.client.Get(c.ctx, key).Result()
		if err != nil {
			continue // Skip failed keys
		}

		var priceData PriceData
		if err := json.Unmarshal([]byte(data), &priceData); err != nil {
			continue // Skip invalid data
		}

		prices = append(prices, priceData)
	}

	return prices, nil
}

// IsPriceStale checks if the cached price is older than the specified duration
func (c *Cache) IsPriceStale(maxAge time.Duration) (bool, error) {
	priceData, err := c.GetCachedPrice()
	if err != nil {
		return true, err // Consider stale if we can't get the price
	}

	return time.Since(priceData.Timestamp) > maxAge, nil
}

// Close closes the Redis connection
func (c *Cache) Close() error {
	return c.client.Close()
}

// Ping tests the Redis connection
func (c *Cache) Ping() error {
	return c.client.Ping(c.ctx).Err()
}
