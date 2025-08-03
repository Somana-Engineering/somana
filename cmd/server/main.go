package main

import (
	"log"

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
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	// Use CORS middleware
	r.Use(cors.New(config))

	// Create service instance
	hostService := services.NewHostService()

	// Register handlers with the generated server
	generated.RegisterHandlers(r, hostService)

	// Start server
	log.Println("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 