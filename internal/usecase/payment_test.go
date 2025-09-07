package usecase

import (
	"payment-service/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPaymentRepository is a mock implementation of PaymentRepository
type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Store(payment *entity.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) GetByTransactionID(transactionID string) (*entity.Payment, error) {
	args := m.Called(transactionID)
	return args.Get(0).(*entity.Payment), args.Error(1)
}

func (m *MockPaymentRepository) Exists(transactionID string) bool {
	args := m.Called(transactionID)
	return args.Bool(0)
}

func TestPaymentUseCase_ProcessPayment_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockPaymentRepository)
	useCase := NewPaymentUseCase(mockRepo)

	req := PaymentRequest{
		UserID:        "user123",
		Amount:        100.50,
		TransactionID: "txn123",
	}

	mockRepo.On("Exists", "txn123").Return(false)
	mockRepo.On("Store", mock.AnythingOfType("*entity.Payment")).Return(nil)

	// Act
	response, err := useCase.ProcessPayment(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "txn123", response.TransactionID)
	assert.Equal(t, "user123", response.UserID)
	assert.Equal(t, 100.50, response.Amount)
	assert.Equal(t, entity.StatusCompleted, response.Status)
	assert.Equal(t, "Payment processed successfully", response.Message)

	mockRepo.AssertExpectations(t)
}

func TestPaymentUseCase_ProcessPayment_IdempotentRequest(t *testing.T) {
	// Arrange
	mockRepo := new(MockPaymentRepository)
	useCase := NewPaymentUseCase(mockRepo)

	req := PaymentRequest{
		UserID:        "user123",
		Amount:        100.50,
		TransactionID: "txn123",
	}

	existingPayment := &entity.Payment{
		TransactionID: "txn123",
		UserID:        "user123",
		Amount:        100.50,
		Status:        entity.StatusCompleted,
	}

	mockRepo.On("Exists", "txn123").Return(true)
	mockRepo.On("GetByTransactionID", "txn123").Return(existingPayment, nil)

	// Act
	response, err := useCase.ProcessPayment(req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "txn123", response.TransactionID)
	assert.Equal(t, "user123", response.UserID)
	assert.Equal(t, 100.50, response.Amount)
	assert.Equal(t, entity.StatusCompleted, response.Status)
	assert.Equal(t, "Transaction already processed", response.Message)

	mockRepo.AssertExpectations(t)
}

func TestPaymentUseCase_ProcessPayment_ValidationErrors(t *testing.T) {
	// Arrange
	mockRepo := new(MockPaymentRepository)
	useCase := NewPaymentUseCase(mockRepo)

	testCases := []struct {
		name        string
		request     PaymentRequest
		expectedErr error
	}{
		{
			name: "Invalid UserID",
			request: PaymentRequest{
				UserID:        "",
				Amount:        100.50,
				TransactionID: "txn123",
			},
			expectedErr: ErrInvalidUserID,
		},
		{
			name: "Invalid TransactionID",
			request: PaymentRequest{
				UserID:        "user123",
				Amount:        100.50,
				TransactionID: "",
			},
			expectedErr: ErrInvalidTransaction,
		},
		{
			name: "Invalid Amount - Zero",
			request: PaymentRequest{
				UserID:        "user123",
				Amount:        0,
				TransactionID: "txn123",
			},
			expectedErr: ErrInvalidAmount,
		},
		{
			name: "Invalid Amount - Negative",
			request: PaymentRequest{
				UserID:        "user123",
				Amount:        -10.50,
				TransactionID: "txn123",
			},
			expectedErr: ErrInvalidAmount,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			response, err := useCase.ProcessPayment(tc.request)

			// Assert
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr, err)
			assert.NotNil(t, response)
			assert.Equal(t, entity.StatusFailed, response.Status)
			assert.Equal(t, tc.expectedErr.Error(), response.Message)
		})
	}
}
