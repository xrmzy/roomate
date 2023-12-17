package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserRepository struct {
	mock.Mock
}

type UserUseCaseTestSuite struct {
	suite.Suite
	uut *usecasemock.MockUserUseCase
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.uut = &usecasemock.MockUserUseCase{}
}

func (m *MockUserRepository) GetAll(limit, offset int) ([]entity.User, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) Get(id string) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Create(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(id string, user entity.User) (entity.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(id, password string) (entity.User, error) {
	args := m.Called(id, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := usecase.NewUserUseCase(mockRepo, nil)

	// Prepare mock data
	mockUsers := []entity.User{
		{Id: "1", Name: "Alice"},
		{Id: "2", Name: "Bob"},
	}

	// Set the expectations
	mockRepo.On("GetAll", 10, 0).Return(mockUsers, nil)

	// Test GetAllUsers method
	users, err := useCase.GetAllUsers(dto.GetAllParams{Limit: 10, Offset: 0})
	assert.NoError(t, err)
	assert.Equal(t, len(mockUsers), len(users))
	// Add more assertions as needed

	// Verify that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_ErrorCase(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := usecase.NewUserUseCase(mockRepo, nil)

	// Prepare mock error
	mockErr := errors.New("repository error")

	// Set the expectations for the mocked GetAll method to return an error
	mockRepo.On("GetAll", 10, 0).Return([]entity.User{}, mockErr) // Return empty slice and error

	// Test GetAllUsers method
	users, err := useCase.GetAllUsers(dto.GetAllParams{Limit: 10, Offset: 0})

	// Ensure that an error is returned
	assert.Error(t, err)
	assert.Nil(t, users) // Ensure that users is nil when an error occurs

	// Verify that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := usecase.NewUserUseCase(mockRepo, nil)

	// Prepare mock user and no error
	mockUser := entity.User{Id: "123", Name: "John Doe"}
	mockRepo.On("Get", "123").Return(mockUser, nil)

	// Test GetUser method
	user, err := useCase.GetUser("123")

	// Ensure no error and user data is correct
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// Verify that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetUser_ErrorCase(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := usecase.NewUserUseCase(mockRepo, nil)

	// Prepare mock error
	mockErr := errors.New("user not found error")

	// Set the expectations for the mocked Get method to return an error
	mockRepo.On("Get", "123").Return(entity.User{}, mockErr) // Return empty user and error

	// Test GetUser method
	user, err := useCase.GetUser("123")

	// Ensure an error is returned and user remains unchanged
	expectedError := fmt.Errorf("user with ID 123 not found")
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, user) // User should be empty

	// Verify that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_SuccessCase(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := usecase.NewUserUseCase(mockRepo, nil) // Menggunakan nil karena tidak menggunakan RoleUseCase

	// Prepare mock data
	mockUserID := "123"
	mockRepo.On("Delete", mockUserID).Return(nil)

	// Test DeleteUser method
	err := useCase.DeleteUser(mockUserID)

	// Ensure no error
	assert.NoError(t, err)

	// Verify that the expectations were met
	mockRepo.AssertExpectations(t)
}
