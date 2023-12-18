package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
)

type CustomerUseCase interface {
	GetAllCustomers(payload dto.GetAllParams) ([]entity.Customer, error)
	GetCustomer(id string) (entity.Customer, error)
	CreateCustomer(customer entity.Customer) (entity.Customer, error)
	UpdateCustomer(id string, customer entity.Customer) (entity.Customer, error)
	DeleteCustomer(id string) error
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (u *customerUseCase) GetAllCustomers(payload dto.GetAllParams) ([]entity.Customer, error) {
	customers, err := u.customerRepo.GetAll(payload.Limit, payload.Offset)

	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (u *customerUseCase) GetCustomer(id string) (entity.Customer, error) {
	customer, err := u.customerRepo.Get(id)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (u *customerUseCase) CreateCustomer(customer entity.Customer) (entity.Customer, error) {
	customer, err := u.customerRepo.Create(customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (u *customerUseCase) UpdateCustomer(id string, customer entity.Customer) (entity.Customer, error) {
	customer, err := u.customerRepo.Update(id, customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (u *customerUseCase) DeleteCustomer(id string) error {
	err := u.customerRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func NewCustomerUseCase(customerRepo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
	}
}
