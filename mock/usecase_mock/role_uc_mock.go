package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type RoleUseCaseMock struct {
	mock.Mock
}

func (r *RoleUseCaseMock) GetRole(id string) (entity.Role, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (r *RoleUseCaseMock) CreateRole(role entity.Role) (entity.Role, error) {
	args := r.Called(role)
	return args.Get(0).(entity.Role), args.Error(1)
}

func (r *RoleUseCaseMock) DeleteRole(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *RoleUseCaseMock) GetAllRoles(payload dto.GetAllParams) ([]entity.Role, error) {
	args := r.Called(payload)
	return args.Get(0).([]entity.Role), args.Error(1)
}

func (r *RoleUseCaseMock) UpdateRole(id string, role entity.Role) (entity.Role, error) {
	args := r.Called(id, role)
	return args.Get(0).(entity.Role), args.Error(1)
}
