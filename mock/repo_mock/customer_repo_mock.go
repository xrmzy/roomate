package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type CustomerRepoMock struct {
	mock.Mock
}

func (r *CustomerRepoMock) Get(id string) (entity.Customer, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (r *CustomerRepoMock) GetAll(limit, offset int) ([]entity.Customer, error) {
	args := r.Called(limit, offset)
	return args.Get(0).([]entity.Customer), args.Error(1)
}

func (r *CustomerRepoMock) Create(customer entity.Customer) (entity.Customer, error) {
	args := r.Called(customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (r *CustomerRepoMock) Update(id string, customer entity.Customer) (entity.Customer, error) {
	args := r.Called(id, customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (r *CustomerRepoMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
