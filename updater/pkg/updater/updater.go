package updater

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/114windd/DeFiOraclePipeline.git/updater/pkg/ethclient"
	"github.com/nats-io/nats.go"
)

// PriceMessage represents a price message from NATS
type PriceMessage struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	ID        string    `json:"id"`
}

// Updater handles consuming price updates and submitting them to the blockchain
type Updater struct {
	conn      *nats.Conn
	ethClient *ethclient.EthClient
	subject   string
	lastPrice float64
	threshold float64
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewUpdater creates a new updater instance
func NewUpdater(natsURL, subject string, ethClient *ethclient.EthClient, threshold float64) (*Updater, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Updater{
		conn:      conn,
		ethClient: ethClient,
		subject:   subject,
		threshold: threshold,
		ctx:       ctx,
		cancel:    cancel,
	}, nil
}

// Start begins consuming price updates from NATS
func (u *Updater) Start() error {
	log.Printf("Starting updater worker for subject: %s", u.subject)

	// Subscribe to price updates
	sub, err := u.conn.Subscribe(u.subject, u.handlePriceMessage)
	if err != nil {
		return fmt.Errorf("failed to subscribe to %s: %w", u.subject, err)
	}
	defer sub.Unsubscribe()

	// Wait for context cancellation
	<-u.ctx.Done()
	log.Println("Updater worker stopped")
	return nil
}

// Stop stops the updater worker
func (u *Updater) Stop() {
	u.cancel()
	u.conn.Close()
}

// handlePriceMessage processes incoming price messages
func (u *Updater) handlePriceMessage(m *nats.Msg) {
	var priceMsg PriceMessage
	if err := json.Unmarshal(m.Data, &priceMsg); err != nil {
		log.Printf("Failed to unmarshal price message: %v", err)
		return
	}

	log.Printf("Received price update: $%.2f from %s", priceMsg.Price, priceMsg.Source)

	// Filter price update
	if !u.FilterPriceUpdate(priceMsg.Price, u.lastPrice) {
		log.Printf("Price change below threshold, skipping update")
		return
	}

	// Validate price
	if err := u.validatePrice(priceMsg.Price); err != nil {
		log.Printf("Price validation failed: %v", err)
		return
	}

	// Submit to blockchain
	txHash, err := u.SendPriceOnChain(priceMsg.Price)
	if err != nil {
		log.Printf("Failed to send price on-chain: %v", err)
		// Retry with exponential backoff
		go u.retryFailedTx(priceMsg.Price, 1)
		return
	}

	log.Printf("Successfully submitted price update. TX: %s", txHash)
	u.lastPrice = priceMsg.Price
}

// FilterPriceUpdate checks if the new price should be pushed (e.g., >0.5% change)
func (u *Updater) FilterPriceUpdate(newPrice, lastPrice float64) bool {
	if lastPrice == 0 {
		return true // Always update if no previous price
	}

	change := (newPrice - lastPrice) / lastPrice
	if change < 0 {
		change = -change // Make it positive
	}

	return change >= u.threshold
}

// validatePrice checks if the price is within reasonable bounds
func (u *Updater) validatePrice(price float64) error {
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

	// Check if price is stale (older than 5 minutes)
	// Note: This would need timestamp validation in a real implementation

	return nil
}

// SendPriceOnChain submits a transaction to update the price on the Solidity contract
func (u *Updater) SendPriceOnChain(price float64) (string, error) {
	if u.ethClient == nil {
		return "", fmt.Errorf("Ethereum client not initialized")
	}

	// Convert price to integer with 8 decimal precision
	priceInt := int64(price * 100000000) // 8 decimal places

	txHash, err := u.ethClient.UpdatePrice(priceInt)
	if err != nil {
		return "", fmt.Errorf("failed to update price on-chain: %w", err)
	}

	return txHash, nil
}

// retryFailedTx retries failed Ethereum transactions with exponential backoff
func (u *Updater) retryFailedTx(price float64, attempt int) {
	if attempt > 5 { // Max 5 retries
		log.Printf("Max retries reached for price update: $%.2f", price)
		return
	}

	// Exponential backoff: 2^attempt seconds
	delay := time.Duration(1<<uint(attempt)) * time.Second
	log.Printf("Retrying price update in %v (attempt %d)", delay, attempt)

	time.Sleep(delay)

	txHash, err := u.SendPriceOnChain(price)
	if err != nil {
		log.Printf("Retry %d failed: %v", attempt, err)
		go u.retryFailedTx(price, attempt+1)
		return
	}

	log.Printf("Retry %d successful. TX: %s", attempt, txHash)
	u.lastPrice = price
}

// GetLastPrice returns the last processed price
func (u *Updater) GetLastPrice() float64 {
	return u.lastPrice
}

// SetThreshold sets the price change threshold
func (u *Updater) SetThreshold(threshold float64) {
	u.threshold = threshold
}

// GetThreshold returns the current threshold
func (u *Updater) GetThreshold() float64 {
	return u.threshold
}

