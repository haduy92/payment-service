package handler

import (
	"encoding/json"
	"net/http"
	"payment-service/internal/usecase"

	"github.com/go-chi/chi/v5"
)

// PaymentHandler handles HTTP requests for payments
type PaymentHandler struct {
	paymentUseCase usecase.PaymentUseCaseInterface
}

// NewPaymentHandler creates a new payment handler
func NewPaymentHandler(paymentUseCase usecase.PaymentUseCaseInterface) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

// ProcessPayment handles POST /pay requests
// @Summary Process Payment
// @Description Processes a payment request with idempotency support. Retrying the same transaction_id will not charge twice.
// @Tags Payments
// @Accept json
// @Produce json
// @Param payment body usecase.PaymentRequest true "Payment request"
// @Success 200 {object} usecase.PaymentResponse "Payment processed successfully"
// @Failure 400 {object} usecase.PaymentResponse "Bad request - validation error"
// @Failure 500 {object} usecase.PaymentResponse "Internal server error"
// @Router /pay [post]
func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req usecase.PaymentRequest

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Process payment through use case
	response, err := h.paymentUseCase.ProcessPayment(req)
	if err != nil {
		// Check if it's a validation error
		switch err {
		case usecase.ErrInvalidAmount, usecase.ErrInvalidUserID, usecase.ErrInvalidTransaction:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Set content type and encode response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetupRoutes configures the HTTP routes
func (h *PaymentHandler) SetupRoutes() chi.Router {
	r := chi.NewRouter()

	// Add middleware for JSON content type
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Post("/pay", h.ProcessPayment)

	return r
}
