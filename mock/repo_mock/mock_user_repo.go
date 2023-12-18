package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll(limit, offset int) ([]entity.User, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) Get(id string) (entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

// Implement other methods of UserRepository interface similarly
