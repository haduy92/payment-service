package repository

import (
	"payment-service/internal/entity"
	"sync"
)

// InMemoryPaymentRepository implements PaymentRepository using in-memory storage
type InMemoryPaymentRepository struct {
	payments map[string]*entity.Payment
	mutex    sync.RWMutex
}

// NewInMemoryPaymentRepository creates a new in-memory payment repository
func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		payments: make(map[string]*entity.Payment),
		mutex:    sync.RWMutex{},
	}
}

// Store saves a payment to the in-memory storage
func (r *InMemoryPaymentRepository) Store(payment *entity.Payment) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.payments[payment.TransactionID] = payment
	return nil
}

// GetByTransactionID retrieves a payment by transaction ID
func (r *InMemoryPaymentRepository) GetByTransactionID(transactionID string) (*entity.Payment, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	payment, exists := r.payments[transactionID]
	if !exists {
		return nil, nil
	}

	return payment, nil
}

// Exists checks if a payment with the given transaction ID exists
func (r *InMemoryPaymentRepository) Exists(transactionID string) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, exists := r.payments[transactionID]
	return exists
}
