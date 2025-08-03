package services

// ExampleService provides example functionality
type ExampleService struct{}

// NewExampleService creates a new example service
func NewExampleService() *ExampleService {
	return &ExampleService{}
}

// GetMessage returns a hardcoded message
func (s *ExampleService) GetMessage() string {
	return "Hello from Somana API!"
} 