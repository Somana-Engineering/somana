package services

import (
	"testing"
)

func TestExampleService_GetMessage(t *testing.T) {
	service := NewExampleService()
	
	message := service.GetMessage()
	
	expected := "Hello from Somana API!"
	if message != expected {
		t.Errorf("Expected message '%s', got '%s'", expected, message)
	}
} 