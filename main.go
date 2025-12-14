package main

import (
	"auth-api/internal/config"
	"auth-api/internal/handlers"
	"auth-api/internal/router"
	"log"

	_ "auth-api/docs" // swagger docs
)

// @title Auth API
// @version 1.0.0
// @description Authentication and authorization API for OAuth2/OIDC flows, token management, MFA/2FA, and SSO integration
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /api/v1

// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the access token.

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()
	authHandler := handlers.NewAuthHandler(cfg)

	// Setup routes
	appRouter := router.NewRouter()
	router.RegisterHealthRoutes(appRouter, healthHandler)
	router.RegisterAuthRoutes(appRouter, authHandler)
	router.RegisterSwaggerRoutes(appRouter)

	// Start server
	server := router.NewServer(cfg, appRouter)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
