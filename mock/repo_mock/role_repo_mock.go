package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type RoleRepoMock struct {
	mock.Mock
}

func (r *RoleRepoMock) Get(id string) (entity.Role, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (r *RoleRepoMock) GetAll(limit, offset int) ([]entity.Role, error) {
	args := r.Called(limit, offset)
	return args.Get(0).([]entity.Role), args.Error(1)
}

func (r *RoleRepoMock) Create(role entity.Role) (entity.Role, error) {
	args := r.Called(role)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (r *RoleRepoMock) Update(id string, role entity.Role) (entity.Role, error) {
	args := r.Called(id, role)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (r *RoleRepoMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
