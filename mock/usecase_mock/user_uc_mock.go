package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) GetAllUsers(payload dto.GetAllParams) ([]entity.User, error) {
	args := u.Called(payload)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (u *UserUseCaseMock) GetUser(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) GetByEmailPassword(email, password string) (entity.User, error) {
	args := u.Called(email, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) CreateUser(user entity.User) (entity.User, error) {
	args := u.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) UpdateUser(id string, user entity.User) (entity.User, error) {
	args := u.Called(id, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) UpdatePassword(id, password string) (entity.User, error) {
	args := u.Called(id, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) DeleteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
