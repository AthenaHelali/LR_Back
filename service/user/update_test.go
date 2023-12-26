package user

import (
	"game-app/param"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockRepo is a mock implementation of your repository interface
func TestUpdateUser(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockRepo)

	// Create a service instance with the mock repository
	service := Service{repo: mockRepo}

	// Define test data
	req := param.UpdateUserRequest{
		Name:        "athena",
		PhoneNumber: "09198888888",
		Password:    "1234567891",
	}

	// Hash the password for comparison
	pass := []byte(req.Password)
	hashedPass, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	req.Password = string(hashedPass)

	// Expect the UpdateUser method to be called with the specified user argument
	mockRepo.On("UpdateUser", mock.AnythingOfType("entity.User")).Return(nil)

	// Call the actual function
	response, err := service.UpdateUser(req)

	// Assert that the mock was called as expected
	mockRepo.AssertExpectations(t)

	// Assert the response
	assert.True(t, response.Success)
	assert.NoError(t, err)
}
