package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukewing/somana/internal/services"
)

// ExampleHandler handles example requests
type ExampleHandler struct {
	exampleService *services.ExampleService
}

// NewExampleHandler creates a new example handler
func NewExampleHandler(exampleService *services.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		exampleService: exampleService,
	}
}

// GetMessage returns a hardcoded message from the service
// @Summary Get example message
// @Description Get a hardcoded message from the example service
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/example [get]
func (h *ExampleHandler) GetMessage(c *gin.Context) {
	message := h.exampleService.GetMessage()
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
} 