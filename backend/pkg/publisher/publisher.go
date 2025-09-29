package publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// PriceMessage represents a message published to NATS
type PriceMessage struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	ID        string    `json:"id"`
}

// Publisher handles publishing price updates to NATS
type Publisher struct {
	conn    *nats.Conn
	subject string
}

// NewPublisher creates a new publisher instance
func NewPublisher(natsURL, subject string) (*Publisher, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &Publisher{
		conn:    conn,
		subject: subject,
	}, nil
}

// NewPublisherWithConn creates a publisher with an existing NATS connection
func NewPublisherWithConn(conn *nats.Conn, subject string) *Publisher {
	return &Publisher{
		conn:    conn,
		subject: subject,
	}
}

// PublishPrice publishes a normalized ETH/USD price to the NATS topic
func (p *Publisher) PublishPrice(price float64, timestamp time.Time, source string) error {
	message := PriceMessage{
		Price:     price,
		Timestamp: timestamp,
		Source:    source,
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal price message: %w", err)
	}

	if err := p.conn.Publish(p.subject, data); err != nil {
		return fmt.Errorf("failed to publish price: %w", err)
	}

	return nil
}

// PublishPriceWithFilter publishes a price only if it meets certain criteria
func (p *Publisher) PublishPriceWithFilter(price float64, timestamp time.Time, source string, lastPrice float64, threshold float64) error {
	// Calculate percentage change
	if lastPrice > 0 {
		change := (price - lastPrice) / lastPrice
		if change < 0 {
			change = -change // Make it positive for comparison
		}

		// Only publish if change is above threshold
		if change < threshold {
			return nil // Skip publishing
		}
	}

	return p.PublishPrice(price, timestamp, source)
}

// PublishPriceAsync publishes a price asynchronously
func (p *Publisher) PublishPriceAsync(price float64, timestamp time.Time, source string) error {
	message := PriceMessage{
		Price:     price,
		Timestamp: timestamp,
		Source:    source,
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal price message: %w", err)
	}

	// Use Publish for async operation (NATS is inherently async)
	err = p.conn.Publish(p.subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish price async: %w", err)
	}

	return nil
}

// PublishBatch publishes multiple prices in a batch
func (p *Publisher) PublishBatch(prices []PriceMessage) error {
	for _, message := range prices {
		data, err := json.Marshal(message)
		if err != nil {
			return fmt.Errorf("failed to marshal price message: %w", err)
		}

		if err := p.conn.Publish(p.subject, data); err != nil {
			return fmt.Errorf("failed to publish price in batch: %w", err)
		}
	}

	return nil
}

// Close closes the NATS connection
func (p *Publisher) Close() {
	p.conn.Close()
}

// IsConnected checks if the NATS connection is active
func (p *Publisher) IsConnected() bool {
	return p.conn.IsConnected()
}

// GetSubject returns the current subject
func (p *Publisher) GetSubject() string {
	return p.subject
}

// SetSubject changes the subject for publishing
func (p *Publisher) SetSubject(subject string) {
	p.subject = subject
}
