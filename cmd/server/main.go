package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lukewing/somana/internal/database"
	"github.com/lukewing/somana/internal/generated"
	"github.com/lukewing/somana/internal/services"
)

func main() {
	// Connect to database
	database.Connect()

	// Create Gin router
	r := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	// Allow all origins in production, or specific origins in development
	if os.Getenv("GIN_MODE") == "release" {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = []string{"http://localhost:3000"}
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	// Use CORS middleware
	r.Use(cors.New(config))

	// Create service instance
	hostService := services.NewHostService()

	// Register handlers with the generated server
	generated.RegisterHandlers(r, hostService)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 