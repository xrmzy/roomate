package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) Get(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) GetAll(limit, offset int) ([]entity.User, error) {
	args := u.Called(limit, offset)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (u *UserRepoMock) GetByEmail(email string) (entity.User, error) {
	args := u.Called(email)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) Create(user entity.User) (entity.User, error) {
	args := u.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) Update(id string, user entity.User) (entity.User, error) {
	args := u.Called(id, user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) UpdatePassword(id, password string) (entity.User, error) {
	args := u.Called(id, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) Delete(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
