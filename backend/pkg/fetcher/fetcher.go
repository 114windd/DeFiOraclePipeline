package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PriceResponse represents the response structure from CoinGecko API
type PriceResponse struct {
	Ethereum struct {
		USD float64 `json:"usd"`
	} `json:"ethereum"`
}

// Fetcher handles fetching ETH/USD prices from external APIs
type Fetcher struct {
	client  *http.Client
	apiURL  string
	timeout time.Duration
}

// NewFetcher creates a new fetcher instance
func NewFetcher(apiURL string, timeout time.Duration) *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: timeout,
		},
		apiURL:  apiURL,
		timeout: timeout,
	}
}

// FetchETHUSDPrice fetches the latest ETH/USD price from the configured API
func (f *Fetcher) FetchETHUSDPrice() (float64, error) {
	req, err := http.NewRequest("GET", f.apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers for better API compatibility
	req.Header.Set("User-Agent", "DeFiOraclePipeline/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := f.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch price: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	var priceResp PriceResponse
	if err := json.Unmarshal(body, &priceResp); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if priceResp.Ethereum.USD <= 0 {
		return 0, fmt.Errorf("invalid price received: %f", priceResp.Ethereum.USD)
	}

	return priceResp.Ethereum.USD, nil
}

// FetchMultiplePrices fetches prices from multiple sources and aggregates them
func (f *Fetcher) FetchMultiplePrices(apiURLs []string) ([]float64, error) {
	var prices []float64
	var errors []error

	// Create a channel to collect results
	type result struct {
		price float64
		err   error
	}

	results := make(chan result, len(apiURLs))

	// Fetch from each API concurrently
	for _, url := range apiURLs {
		go func(apiURL string) {
			fetcher := NewFetcher(apiURL, f.timeout)
			price, err := fetcher.FetchETHUSDPrice()
			results <- result{price: price, err: err}
		}(url)
	}

	// Collect results
	for i := 0; i < len(apiURLs); i++ {
		res := <-results
		if res.err != nil {
			errors = append(errors, res.err)
			continue
		}
		prices = append(prices, res.price)
	}

	if len(prices) == 0 {
		return nil, fmt.Errorf("failed to fetch from any API: %v", errors)
	}

	return prices, nil
}

// AggregatePrices combines multiple exchange prices into a single normalized value
// Uses simple average for now, but could be enhanced with weighted averages
func AggregatePrices(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}

	var sum float64
	for _, price := range prices {
		sum += price
	}

	return sum / float64(len(prices))
}
