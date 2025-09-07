package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"payment-service/internal/entity"
	"payment-service/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPaymentUseCase is a mock implementation of PaymentUseCaseInterface
type MockPaymentUseCase struct {
	mock.Mock
}

func (m *MockPaymentUseCase) ProcessPayment(req usecase.PaymentRequest) (*usecase.PaymentResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*usecase.PaymentResponse), args.Error(1)
}

func TestPaymentHandler_ProcessPayment_Success(t *testing.T) {
	// Arrange
	mockUseCase := new(MockPaymentUseCase)
	handler := NewPaymentHandler(mockUseCase)

	requestBody := usecase.PaymentRequest{
		UserID:        "user123",
		Amount:        100.50,
		TransactionID: "txn123",
	}

	expectedResponse := &usecase.PaymentResponse{
		TransactionID: "txn123",
		UserID:        "user123",
		Amount:        100.50,
		Status:        entity.StatusCompleted,
		Message:       "Payment processed successfully",
	}

	mockUseCase.On("ProcessPayment", requestBody).Return(expectedResponse, nil)

	// Create request
	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/pay", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Act
	handler.ProcessPayment(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response usecase.PaymentResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.TransactionID, response.TransactionID)
	assert.Equal(t, expectedResponse.UserID, response.UserID)
	assert.Equal(t, expectedResponse.Amount, response.Amount)
	assert.Equal(t, expectedResponse.Status, response.Status)

	mockUseCase.AssertExpectations(t)
}

func TestPaymentHandler_ProcessPayment_ValidationError(t *testing.T) {
	// Arrange
	mockUseCase := new(MockPaymentUseCase)
	handler := NewPaymentHandler(mockUseCase)

	requestBody := usecase.PaymentRequest{
		UserID:        "",
		Amount:        100.50,
		TransactionID: "txn123",
	}

	expectedResponse := &usecase.PaymentResponse{
		TransactionID: "txn123",
		UserID:        "",
		Amount:        100.50,
		Status:        entity.StatusFailed,
		Message:       "user ID cannot be empty",
	}

	mockUseCase.On("ProcessPayment", requestBody).Return(expectedResponse, usecase.ErrInvalidUserID)

	// Create request
	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/pay", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	rr := httptest.NewRecorder()

	// Act
	handler.ProcessPayment(rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response usecase.PaymentResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, entity.StatusFailed, response.Status)
	assert.Equal(t, "user ID cannot be empty", response.Message)

	mockUseCase.AssertExpectations(t)
}
