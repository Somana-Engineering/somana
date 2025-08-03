package main

import (
	"log"

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