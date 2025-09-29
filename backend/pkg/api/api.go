package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/cache"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/metrics"
	"github.com/114windd/DeFiOraclePipeline.git/backend/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// API handles HTTP endpoints
type API struct {
	router  *gin.Engine
	cache   *cache.Cache
	storage *storage.Storage
	metrics *metrics.Metrics
}

// NewAPI creates a new API instance
func NewAPI(cache *cache.Cache, storage *storage.Storage, metrics *metrics.Metrics) *API {
	router := gin.Default()

	api := &API{
		router:  router,
		cache:   cache,
		storage: storage,
		metrics: metrics,
	}

	api.setupRoutes()
	return api
}

// setupRoutes configures all the API routes
func (a *API) setupRoutes() {
	// Health check endpoint
	a.router.GET("/health", a.healthCheck)

	// Price endpoints
	a.router.GET("/price", a.getLatestPrice)
	a.router.GET("/price/history", a.getPriceHistory)
	a.router.GET("/price/twap", a.getTWAP)

	// Metrics endpoint
	a.router.GET("/metrics", a.getMetrics)

	// Admin endpoints
	a.router.GET("/admin/stats", a.getStats)
}

// healthCheck returns the health status of the service
func (a *API) healthCheck(c *gin.Context) {
	// Check cache connection
	if err := a.cache.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "cache connection failed",
		})
		return
	}

	// Check database connection
	if err := a.storage.GetDB().Error; err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  "database connection failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
	})
}

// getLatestPrice returns the latest ETH/USD price
func (a *API) getLatestPrice(c *gin.Context) {
	start := time.Now()

	// Try to get from cache first
	priceData, err := a.cache.GetCachedPrice()
	if err != nil {
		// Fallback to database
		record, dbErr := a.storage.GetLatestPrice()
		if dbErr != nil {
			a.metrics.RecordCacheError("redis", "get")
			a.metrics.RecordDBError("select", "price_records", "not_found")
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no price data available",
			})
			return
		}

		priceData = &cache.PriceData{
			Price:     record.Price,
			Timestamp: record.Timestamp,
			Source:    record.Source,
		}
		a.metrics.RecordCacheMiss("redis")
	} else {
		a.metrics.RecordCacheHit("redis")
	}

	// Record metrics
	a.metrics.RecordPriceAge(time.Since(priceData.Timestamp), priceData.Source)

	c.JSON(http.StatusOK, gin.H{
		"price":       priceData.Price,
		"timestamp":   priceData.Timestamp.Unix(),
		"source":      priceData.Source,
		"age_seconds": time.Since(priceData.Timestamp).Seconds(),
	})

	// Record latency
	a.metrics.RecordFetchLatency(time.Since(start), "api", "success")
}

// getPriceHistory returns historical price data
func (a *API) getPriceHistory(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 1000 {
		limit = 100
	}

	// Try cache first
	history, err := a.cache.GetPriceHistory(limit)
	if err != nil || len(history) == 0 {
		// Fallback to database
		records, dbErr := a.storage.GetPriceHistory(limit)
		if dbErr != nil {
			a.metrics.RecordDBError("select", "price_records", "query_failed")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to retrieve price history",
			})
			return
		}

		// Convert records to cache format
		history = make([]cache.PriceData, len(records))
		for i, record := range records {
			history[i] = cache.PriceData{
				Price:     record.Price,
				Timestamp: record.Timestamp,
				Source:    record.Source,
			}
		}
		a.metrics.RecordCacheMiss("redis")
	} else {
		a.metrics.RecordCacheHit("redis")
	}

	c.JSON(http.StatusOK, gin.H{
		"prices": history,
		"count":  len(history),
	})
}

// getTWAP returns the Time-Weighted Average Price
func (a *API) getTWAP(c *gin.Context) {
	durationStr := c.DefaultQuery("duration", "1h")
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid duration format. Use formats like '1h', '30m', '24h'",
		})
		return
	}

	start := time.Now()
	twap, err := a.storage.CalculateTWAP(duration)
	if err != nil {
		a.metrics.RecordDBError("select", "price_records", "twap_calculation")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to calculate TWAP",
		})
		return
	}

	a.metrics.RecordDBLatency(time.Since(start), "calculate_twap", "price_records")

	c.JSON(http.StatusOK, gin.H{
		"twap":          twap,
		"duration":      duration.String(),
		"calculated_at": time.Now().Unix(),
	})
}

// getMetrics returns Prometheus metrics
func (a *API) getMetrics(c *gin.Context) {
	// Use the Prometheus handler to return metrics in the correct format
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

// getStats returns system statistics
func (a *API) getStats(c *gin.Context) {
	// Get price count from database
	count, err := a.storage.GetPriceCount()
	if err != nil {
		a.metrics.RecordDBError("count", "price_records", "query_failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get statistics",
		})
		return
	}

	// Get latest price
	latestPrice, err := a.storage.GetLatestPrice()
	if err != nil {
		latestPrice = &storage.PriceRecord{}
	}

	c.JSON(http.StatusOK, gin.H{
		"total_prices": count,
		"latest_price": gin.H{
			"price":     latestPrice.Price,
			"timestamp": latestPrice.Timestamp.Unix(),
			"source":    latestPrice.Source,
		},
		"uptime": time.Now().Unix(),
	})
}

// Run starts the HTTP server
func (a *API) Run(addr string) error {
	return a.router.Run(addr)
}

// GetRouter returns the Gin router for testing
func (a *API) GetRouter() *gin.Engine {
	return a.router
}
