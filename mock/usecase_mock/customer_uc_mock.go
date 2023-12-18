package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type CustomerUseCaseMock struct {
	mock.Mock
}

func (c *CustomerUseCaseMock) GetAllCustomers(payload dto.GetAllParams) ([]entity.Customer, error) {
	args := c.Called(payload)
	return args.Get(0).([]entity.Customer), args.Error(1)
}

func (c *CustomerUseCaseMock) GetCustomer(id string) (entity.Customer, error) {
	args := c.Called(id)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (c *CustomerUseCaseMock) CreateCustomer(customer entity.Customer) (entity.Customer, error) {
	args := c.Called(customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (c *CustomerUseCaseMock) UpdateCustomer(id string, customer entity.Customer) (entity.Customer, error) {
	args := c.Called(id, customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (c *CustomerUseCaseMock) DeleteCustomer(id string) error {
	args := c.Called(id)
	return args.Error(0)
}
