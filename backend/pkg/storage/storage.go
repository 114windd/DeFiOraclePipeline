package storage

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PriceRecord represents a price record in the database
type PriceRecord struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Price     float64   `gorm:"not null;type:decimal(20,8)" json:"price"`
	Timestamp time.Time `gorm:"not null;index" json:"timestamp"`
	Source    string    `gorm:"size:100" json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Storage handles database operations for price persistence
type Storage struct {
	db *gorm.DB
}

// NewStorage creates a new storage instance
func NewStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&PriceRecord{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Storage{db: db}, nil
}

// NewStorageWithDB creates a storage instance with an existing database connection
func NewStorageWithDB(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

// SavePrice stores a price record in the SQL database
func (s *Storage) SavePrice(price float64, timestamp time.Time, source string) error {
	record := PriceRecord{
		Price:     price,
		Timestamp: timestamp,
		Source:    source,
	}

	if err := s.db.Create(&record).Error; err != nil {
		return fmt.Errorf("failed to save price: %w", err)
	}

	return nil
}

// GetPriceHistory retrieves recent price records for computing TWAP
func (s *Storage) GetPriceHistory(limit int) ([]PriceRecord, error) {
	var records []PriceRecord

	if err := s.db.Order("timestamp DESC").Limit(limit).Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to get price history: %w", err)
	}

	return records, nil
}

// GetLatestPrice retrieves the most recent price record
func (s *Storage) GetLatestPrice() (*PriceRecord, error) {
	var record PriceRecord

	if err := s.db.Order("timestamp DESC").First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no price records found")
		}
		return nil, fmt.Errorf("failed to get latest price: %w", err)
	}

	return &record, nil
}

// GetPricesInRange retrieves prices within a specific time range
func (s *Storage) GetPricesInRange(start, end time.Time) ([]PriceRecord, error) {
	var records []PriceRecord

	if err := s.db.Where("timestamp BETWEEN ? AND ?", start, end).
		Order("timestamp ASC").
		Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to get prices in range: %w", err)
	}

	return records, nil
}

// CalculateTWAP calculates the Time-Weighted Average Price for a given period
func (s *Storage) CalculateTWAP(duration time.Duration) (float64, error) {
	end := time.Now()
	start := end.Add(-duration)

	records, err := s.GetPricesInRange(start, end)
	if err != nil {
		return 0, err
	}

	if len(records) == 0 {
		return 0, fmt.Errorf("no price records found in the specified period")
	}

	if len(records) == 1 {
		return records[0].Price, nil
	}

	var totalWeightedPrice float64
	var totalWeight float64

	for i := 0; i < len(records)-1; i++ {
		current := records[i]
		next := records[i+1]

		// Calculate time weight (duration between this price and the next)
		weight := next.Timestamp.Sub(current.Timestamp).Seconds()

		totalWeightedPrice += current.Price * weight
		totalWeight += weight
	}

	// Add the last record with weight to the end time
	lastRecord := records[len(records)-1]
	lastWeight := end.Sub(lastRecord.Timestamp).Seconds()
	totalWeightedPrice += lastRecord.Price * lastWeight
	totalWeight += lastWeight

	if totalWeight == 0 {
		return 0, fmt.Errorf("invalid time weights")
	}

	return totalWeightedPrice / totalWeight, nil
}

// GetPriceCount returns the total number of price records
func (s *Storage) GetPriceCount() (int64, error) {
	var count int64
	if err := s.db.Model(&PriceRecord{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count price records: %w", err)
	}
	return count, nil
}

// DeleteOldRecords deletes price records older than the specified duration
func (s *Storage) DeleteOldRecords(olderThan time.Duration) error {
	cutoff := time.Now().Add(-olderThan)

	if err := s.db.Where("timestamp < ?", cutoff).Delete(&PriceRecord{}).Error; err != nil {
		return fmt.Errorf("failed to delete old records: %w", err)
	}

	return nil
}

// Close closes the database connection
func (s *Storage) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GetDB returns the underlying database connection
func (s *Storage) GetDB() *gorm.DB {
	return s.db
}

