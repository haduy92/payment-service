package usecase

import (
	"payment-service/internal/entity"
	"time"
)

// PaymentUseCase handles payment business logic
type PaymentUseCase struct {
	repo PaymentRepository
}

// NewPaymentUseCase creates a new payment use case
func NewPaymentUseCase(repo PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		repo: repo,
	}
}

// ProcessPayment processes a payment request with idempotency
func (p *PaymentUseCase) ProcessPayment(req PaymentRequest) (*PaymentResponse, error) {
	// Validate request
	if err := p.validateRequest(req); err != nil {
		return &PaymentResponse{
			TransactionID: req.TransactionID,
			UserID:        req.UserID,
			Amount:        req.Amount,
			Status:        entity.StatusFailed,
			Message:       err.Error(),
		}, err
	}

	// Check if transaction already exists (idempotency)
	if p.repo.Exists(req.TransactionID) {
		existingPayment, err := p.repo.GetByTransactionID(req.TransactionID)
		if err != nil {
			return nil, err
		}

		return &PaymentResponse{
			TransactionID: existingPayment.TransactionID,
			UserID:        existingPayment.UserID,
			Amount:        existingPayment.Amount,
			Status:        existingPayment.Status,
			Message:       "Transaction already processed",
		}, nil
	}

	// Create new payment
	payment := &entity.Payment{
		TransactionID: req.TransactionID,
		UserID:        req.UserID,
		Amount:        req.Amount,
		Status:        entity.StatusCompleted,
		CreatedAt:     time.Now(),
	}

	// Store payment
	if err := p.repo.Store(payment); err != nil {
		return &PaymentResponse{
			TransactionID: req.TransactionID,
			UserID:        req.UserID,
			Amount:        req.Amount,
			Status:        entity.StatusFailed,
			Message:       "Failed to process payment",
		}, err
	}

	return &PaymentResponse{
		TransactionID: payment.TransactionID,
		UserID:        payment.UserID,
		Amount:        payment.Amount,
		Status:        payment.Status,
		Message:       "Payment processed successfully",
	}, nil
}

// validateRequest validates the payment request
func (p *PaymentUseCase) validateRequest(req PaymentRequest) error {
	if req.UserID == "" {
		return ErrInvalidUserID
	}
	if req.TransactionID == "" {
		return ErrInvalidTransaction
	}
	if req.Amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}
