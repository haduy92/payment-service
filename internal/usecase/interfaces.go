package usecase

import (
	"errors"
	"payment-service/internal/entity"
)

// PaymentRepository defines the interface for payment storage
type PaymentRepository interface {
	Store(payment *entity.Payment) error
	GetByTransactionID(transactionID string) (*entity.Payment, error)
	Exists(transactionID string) bool
}

// PaymentUseCaseInterface defines the interface for payment use case
type PaymentUseCaseInterface interface {
	ProcessPayment(req PaymentRequest) (*PaymentResponse, error)
}

// PaymentRequest represents the request payload for payment
type PaymentRequest struct {
	UserID        string  `json:"user_id" example:"user123" validate:"required"`        // User ID for the payment
	Amount        float64 `json:"amount" example:"99.99" validate:"required,gt=0"`      // Payment amount (must be greater than 0)
	TransactionID string  `json:"transaction_id" example:"txn-456" validate:"required"` // Unique transaction ID for idempotency
}

// PaymentResponse represents the response for payment
type PaymentResponse struct {
	TransactionID string  `json:"transaction_id" example:"txn-456"`                 // Transaction ID
	UserID        string  `json:"user_id" example:"user123"`                        // User ID
	Amount        float64 `json:"amount" example:"99.99"`                           // Payment amount
	Status        string  `json:"status" example:"success"`                         // Payment status (success, failed)
	Message       string  `json:"message" example:"Payment processed successfully"` // Status message
}

var (
	ErrInvalidAmount        = errors.New("amount must be greater than 0")
	ErrInvalidUserID        = errors.New("user ID cannot be empty")
	ErrInvalidTransaction   = errors.New("transaction ID cannot be empty")
	ErrDuplicateTransaction = errors.New("transaction already processed")
)
