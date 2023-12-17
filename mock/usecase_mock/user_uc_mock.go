package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) GetAllUsers(payload dto.GetAllParams) ([]entity.User, error) {
	args := m.Called(payload)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserUseCase) GetUser(id string) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserUseCase) GetByEmailPassword(email, password string) (entity.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserUseCase) CreateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserUseCase) UpdateUser(id string, user entity.User) (entity.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserUseCase) UpdatePassword(id, password string) (entity.User, error) {
	args := m.Called(id, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserUseCase) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
