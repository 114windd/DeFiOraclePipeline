package normalizer

import (
	"fmt"
	"math"
)

const (
	// DecimalPrecision defines the number of decimal places for normalized prices
	DecimalPrecision = 8
	// Multiplier is 10^8 for 8 decimal places
	Multiplier = 100000000
)

// Normalizer handles price normalization to ensure consistent decimal precision
type Normalizer struct {
	precision  int
	multiplier float64
}

// NewNormalizer creates a new normalizer with specified precision
func NewNormalizer(precision int) *Normalizer {
	multiplier := math.Pow10(precision)
	return &Normalizer{
		precision:  precision,
		multiplier: multiplier,
	}
}

// NewDefaultNormalizer creates a normalizer with 8 decimal precision
func NewDefaultNormalizer() *Normalizer {
	return NewNormalizer(DecimalPrecision)
}

// NormalizePrice ensures consistent decimal precision by rounding to specified decimal places
func (n *Normalizer) NormalizePrice(price float64) float64 {
	if price <= 0 {
		return 0
	}

	// Round to the specified number of decimal places
	rounded := math.Round(price*n.multiplier) / n.multiplier

	return rounded
}

// NormalizePriceToInt converts a price to an integer representation with specified precision
// This is useful for blockchain operations where we need integer values
func (n *Normalizer) NormalizePriceToInt(price float64) int64 {
	if price <= 0 {
		return 0
	}

	normalized := n.NormalizePrice(price)
	return int64(normalized * n.multiplier)
}

// IntToPrice converts an integer representation back to a float price
func (n *Normalizer) IntToPrice(priceInt int64) float64 {
	return float64(priceInt) / n.multiplier
}

// ValidatePrice checks if a price is within reasonable bounds
func (n *Normalizer) ValidatePrice(price float64) error {
	if price <= 0 {
		return fmt.Errorf("price must be positive, got: %f", price)
	}

	// Check for unreasonably high prices (e.g., > $1M per ETH)
	if price > 1000000 {
		return fmt.Errorf("price seems unreasonably high: %f", price)
	}

	// Check for unreasonably low prices (e.g., < $1 per ETH)
	if price < 1 {
		return fmt.Errorf("price seems unreasonably low: %f", price)
	}

	return nil
}

// GetPrecision returns the current precision setting
func (n *Normalizer) GetPrecision() int {
	return n.precision
}

// GetMultiplier returns the current multiplier
func (n *Normalizer) GetMultiplier() float64 {
	return n.multiplier
}
