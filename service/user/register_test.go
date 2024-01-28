package user

import (
	"game-app/entity"
	"game-app/param"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) RegisterUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRepo) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	args := m.Called(phoneNumber)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRepo) GetUserByID(userID uint) (entity.User, error) {
	args := m.Called(userID)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockRepo) AddFavoriteLaptop(userID int, laptopID int) error {
	args := m.Called(userID, laptopID)
	return args.Error(0)
}

func (m *MockRepo) RemoveFavoriteLaptop(laptopID int, userID int) error {
	args := m.Called(laptopID, userID)
	return args.Error(0)
}

func (m *MockRepo) GetLaptops(userID uint) ([]entity.Laptop, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.Laptop), args.Error(1)
}

func (m *MockRepo) GetLaptopByID(laptopID uint) (entity.Laptop, error) {
	args := m.Called(laptopID)
	return args.Get(0).(entity.Laptop), args.Error(1)
}

func (m *MockRepo) UpdateUser(updatedUser entity.User) error {
	args := m.Called(updatedUser)
	return args.Error(0)
}

func (m *MockRepo) Search(IDs []int) ([]param.LaptopInfo, error) {
	args := m.Called(IDs)
	return args.Get(0).([]param.LaptopInfo), args.Error(1)
}

func (m *MockRepo) AddLaptop(LaptopInfo param.LaptopInfo, UserID uint) (uint, error) {
	args := m.Called(LaptopInfo, UserID)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockRepo) AddSellerLaptop(LaptopID, UserID uint) error {
	return nil
}

func (m *MockRepo) GetSellerLaptops(UserID uint) ([]entity.Laptop, error) {
	return nil, nil
}

func (m *MockRepo) UpdateLaptop(updatedLaptop entity.Laptop) error {
	return nil
}

func (m *MockRepo) RemoveSellerLaptop(LaptopID int, SellerID int) error {
	return nil
}

func TestRegister(t *testing.T) {

	// Create a mock repository
	mockRepo := new(MockRepo)

	// Create a service instance with the mock repository
	service := Service{repo: mockRepo}

	// Define test data
	req := param.RegisterRequest{
		Name:        "athena",
		PhoneNumber: "09198888888",
		Password:    "password123",
	}

	// Hash the password for comparison
	pass := []byte(req.Password)
	hashedPass, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	expectedUser := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(hashedPass),
	}

	// Expect the RegisterUser method to be called with the specified user argument
	mockRepo.On("RegisterUser", mock.AnythingOfType("entity.User")).Return(expectedUser, nil)

	// Call the actual function
	response, err := service.Register(req)

	// Assert that the mock was called as expected
	mockRepo.AssertExpectations(t)

	// Assert the response
	assert.NotNil(t, response.User)
	assert.Equal(t, expectedUser.ID, response.User.ID)
	assert.Equal(t, expectedUser.Name, response.User.Name)
	assert.Equal(t, expectedUser.PhoneNumber, response.User.PhoneNumber)
	assert.NoError(t, err)
}
