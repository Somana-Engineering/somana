package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/lukewing/somana/internal/database"
	"github.com/lukewing/somana/internal/server"
	"github.com/lukewing/somana/internal/services"
	// _ "github.com/lukewing/somana/docs" // Uncomment after running make generate
)

// @title           Somana API
// @version         1.0
// @description     A sample API for the Somana project
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	database.Connect()

	// Initialize services
	generatedService := services.NewGeneratedService()

	// Create router
	r := gin.Default()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Register generated handlers
	server.RegisterHandlers(r, generatedService)

	// Swagger documentation (only if docs are generated)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// OpenAPI JSON endpoint (only if docs are generated)
	r.GET("/openapi.json", func(c *gin.Context) {
		// This will serve the generated OpenAPI spec
		c.File("./docs/swagger.json")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 