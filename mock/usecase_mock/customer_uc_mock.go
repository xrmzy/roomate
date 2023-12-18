package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type MockCustomerUseCase struct {
	mock.Mock
}

func (m *MockCustomerUseCase) GetAllCustomers(payload dto.GetAllParams) ([]entity.Customer, error) {
	args := m.Called(payload)
	return args.Get(0).([]entity.Customer), args.Error(1)
}

func (m *MockCustomerUseCase) GetCustomer(id string) (entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (m *MockCustomerUseCase) CreateCustomer(customer entity.Customer) (entity.Customer, error) {
	args := m.Called(customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (m *MockCustomerUseCase) UpdateCustomer(id string, customer entity.Customer) (entity.Customer, error) {
	args := m.Called(id, customer)
	return args.Get(0).(entity.Customer), args.Error(1)
}

func (m *MockCustomerUseCase) DeleteCustomer(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
