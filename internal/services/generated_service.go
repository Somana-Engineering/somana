package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukewing/somana/internal/database"
	"github.com/lukewing/somana/internal/models"
	"github.com/lukewing/somana/internal/server"
)

// GeneratedService implements the generated ServerInterface
type GeneratedService struct{}

// NewGeneratedService creates a new generated service
func NewGeneratedService() *GeneratedService {
	return &GeneratedService{}
}

// GetHealth implements the health check endpoint
func (s *GeneratedService) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Somana API is running",
		"version": "1.0.0",
	})
}

// GetApiV1Resources implements the list resources endpoint
func (s *GeneratedService) GetApiV1Resources(c *gin.Context, params server.GetApiV1ResourcesParams) {
	var resources []models.Resource
	query := database.DB

	// Apply active filter if provided
	if params.Active != nil {
		query = query.Where("active = ?", *params.Active)
	}

	result := query.Find(&resources)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}

// PostApiV1Resources implements the create resource endpoint
func (s *GeneratedService) PostApiV1Resources(c *gin.Context) {
	var req models.ResourceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource := &models.Resource{
		Name:        req.Name,
		Description: req.Description,
		Active:      true,
	}

	result := database.DB.Create(resource)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, resource)
}

// GetApiV1ResourcesActive implements the active resources endpoint
func (s *GeneratedService) GetApiV1ResourcesActive(c *gin.Context) {
	var resources []models.Resource
	result := database.DB.Where("active = ?", true).Find(&resources)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, resources)
}

// GetApiV1ResourcesId implements the get resource by ID endpoint
func (s *GeneratedService) GetApiV1ResourcesId(c *gin.Context, id int) {
	var resource models.Resource
	result := database.DB.First(&resource, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	c.JSON(http.StatusOK, resource)
}

// PutApiV1ResourcesId implements the update resource endpoint
func (s *GeneratedService) PutApiV1ResourcesId(c *gin.Context, id int) {
	var req models.ResourceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resource models.Resource
	result := database.DB.First(&resource, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	// Update fields if provided
	if req.Name != nil {
		resource.Name = *req.Name
	}
	if req.Description != nil {
		resource.Description = req.Description
	}
	if req.Active != nil {
		resource.Active = *req.Active
	}

	result = database.DB.Save(&resource)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, resource)
}

// DeleteApiV1ResourcesId implements the delete resource endpoint
func (s *GeneratedService) DeleteApiV1ResourcesId(c *gin.Context, id int) {
	var resource models.Resource
	result := database.DB.First(&resource, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		return
	}

	result = database.DB.Delete(&resource)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
} 