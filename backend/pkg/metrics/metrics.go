package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics holds all the Prometheus metrics
type Metrics struct {
	// Price fetch metrics
	FetchLatency prometheus.HistogramVec
	FetchErrors  prometheus.CounterVec
	FetchSuccess prometheus.CounterVec

	// Price update metrics
	PriceUpdates prometheus.CounterVec
	PriceAge     prometheus.GaugeVec

	// Cache metrics
	CacheHits   prometheus.CounterVec
	CacheMisses prometheus.CounterVec
	CacheErrors prometheus.CounterVec

	// Database metrics
	DBOperations prometheus.CounterVec
	DBLatency    prometheus.HistogramVec
	DBErrors     prometheus.CounterVec

	// NATS metrics
	NATSPublished prometheus.CounterVec
	NATSErrors    prometheus.CounterVec

	// System metrics
	ActiveConnections prometheus.Gauge
	MemoryUsage       prometheus.Gauge
}

// NewMetrics creates a new metrics instance
func NewMetrics() *Metrics {
	return &Metrics{
		FetchLatency: *promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "price_fetch_duration_seconds",
				Help:    "Duration of price fetch operations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"source", "status"},
		),
		FetchErrors: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "price_fetch_errors_total",
				Help: "Total number of price fetch errors",
			},
			[]string{"source", "error_type"},
		),
		FetchSuccess: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "price_fetch_success_total",
				Help: "Total number of successful price fetches",
			},
			[]string{"source"},
		),
		PriceUpdates: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "price_updates_total",
				Help: "Total number of price updates",
			},
			[]string{"source", "type"},
		),
		PriceAge: *promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "price_age_seconds",
				Help: "Age of the latest price in seconds",
			},
			[]string{"source"},
		),
		CacheHits: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "cache_hits_total",
				Help: "Total number of cache hits",
			},
			[]string{"cache_type"},
		),
		CacheMisses: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "cache_misses_total",
				Help: "Total number of cache misses",
			},
			[]string{"cache_type"},
		),
		CacheErrors: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "cache_errors_total",
				Help: "Total number of cache errors",
			},
			[]string{"cache_type", "operation"},
		),
		DBOperations: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "database_operations_total",
				Help: "Total number of database operations",
			},
			[]string{"operation", "table"},
		),
		DBLatency: *promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "database_operation_duration_seconds",
				Help:    "Duration of database operations",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"operation", "table"},
		),
		DBErrors: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "database_errors_total",
				Help: "Total number of database errors",
			},
			[]string{"operation", "table", "error_type"},
		),
		NATSPublished: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "nats_published_total",
				Help: "Total number of messages published to NATS",
			},
			[]string{"subject"},
		),
		NATSErrors: *promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "nats_errors_total",
				Help: "Total number of NATS errors",
			},
			[]string{"operation"},
		),
		ActiveConnections: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "active_connections",
				Help: "Number of active connections",
			},
		),
		MemoryUsage: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "memory_usage_bytes",
				Help: "Current memory usage in bytes",
			},
		),
	}
}

// RecordFetchLatency records how long it took to fetch a price
func (m *Metrics) RecordFetchLatency(duration time.Duration, source, status string) {
	m.FetchLatency.WithLabelValues(source, status).Observe(duration.Seconds())
}

// RecordFetchError records a price fetch error
func (m *Metrics) RecordFetchError(source, errorType string) {
	m.FetchErrors.WithLabelValues(source, errorType).Inc()
}

// RecordFetchSuccess records a successful price fetch
func (m *Metrics) RecordFetchSuccess(source string) {
	m.FetchSuccess.WithLabelValues(source).Inc()
}

// RecordPriceUpdate increments counter for successful price updates
func (m *Metrics) RecordPriceUpdate(source, updateType string) {
	m.PriceUpdates.WithLabelValues(source, updateType).Inc()
}

// RecordPriceAge records the age of the latest price
func (m *Metrics) RecordPriceAge(age time.Duration, source string) {
	m.PriceAge.WithLabelValues(source).Set(age.Seconds())
}

// RecordCacheHit records a cache hit
func (m *Metrics) RecordCacheHit(cacheType string) {
	m.CacheHits.WithLabelValues(cacheType).Inc()
}

// RecordCacheMiss records a cache miss
func (m *Metrics) RecordCacheMiss(cacheType string) {
	m.CacheMisses.WithLabelValues(cacheType).Inc()
}

// RecordCacheError records a cache error
func (m *Metrics) RecordCacheError(cacheType, operation string) {
	m.CacheErrors.WithLabelValues(cacheType, operation).Inc()
}

// RecordDBOperation records a database operation
func (m *Metrics) RecordDBOperation(operation, table string) {
	m.DBOperations.WithLabelValues(operation, table).Inc()
}

// RecordDBLatency records database operation latency
func (m *Metrics) RecordDBLatency(duration time.Duration, operation, table string) {
	m.DBLatency.WithLabelValues(operation, table).Observe(duration.Seconds())
}

// RecordDBError records a database error
func (m *Metrics) RecordDBError(operation, table, errorType string) {
	m.DBErrors.WithLabelValues(operation, table, errorType).Inc()
}

// RecordNATSPublished records a NATS publish operation
func (m *Metrics) RecordNATSPublished(subject string) {
	m.NATSPublished.WithLabelValues(subject).Inc()
}

// RecordNATSError records a NATS error
func (m *Metrics) RecordNATSError(operation string) {
	m.NATSErrors.WithLabelValues(operation).Inc()
}

// SetActiveConnections sets the number of active connections
func (m *Metrics) SetActiveConnections(count float64) {
	m.ActiveConnections.Set(count)
}

// SetMemoryUsage sets the current memory usage
func (m *Metrics) SetMemoryUsage(bytes float64) {
	m.MemoryUsage.Set(bytes)
}
