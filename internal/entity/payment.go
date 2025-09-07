package entity

import (
	"time"
)

// Payment represents a payment transaction
type Payment struct {
	TransactionID string    `json:"transaction_id"`
	UserID        string    `json:"user_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

// PaymentStatus constants
const (
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)
