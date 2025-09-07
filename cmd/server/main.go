package main

import (
	"fmt"
	"log"
	"net/http"
	"payment-service/internal/handler"
	"payment-service/internal/repository"
	"payment-service/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "payment-service/docs" // Import generated docs
)

// @title Payment Service API
// @version 1.0
// @description A Go-based payment service implementing clean architecture with idempotent payment processing
// @contact.name Payment Service
// @license.name MIT
// @host localhost:8080
// @BasePath /

func main() {
	// Initialize repository
	paymentRepo := repository.NewInMemoryPaymentRepository()

	// Initialize use case
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo)

	// Initialize handler
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	// Setup router
	r := chi.NewRouter()

	// Add middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Mount payment routes
	r.Mount("/", paymentHandler.SetupRoutes())

	// Health check endpoint
	// @Summary Health Check
	// @Description Returns the health status of the payment service
	// @Tags Health
	// @Produce json
	// @Success 200 {object} map[string]string "Service is healthy"
	// @Router /health [get]
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":"ok","service":"payment-service"}`)
	})

	// Swagger documentation endpoint
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Legacy docs endpoint (redirect to swagger)
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	port := ":8080"
	fmt.Printf("Payment service starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
