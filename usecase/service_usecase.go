package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
	"roomate/utils/common"
)

type ServiceUseCase interface {
	GetAllServices(payload dto.GetAllParams) ([]entity.Service, error)
	GetService(id string) (entity.Service, error)
	CreateService(service entity.Service) (entity.Service, error)
	UpdateService(id string, service entity.Service) (entity.Service, error)
	DeleteService(id string) error
}

type serviceUseCase struct {
	serviceRepo repository.ServiceRepository
}

func (u *serviceUseCase) GetAllServices(payload dto.GetAllParams) ([]entity.Service, error) {
	services, err := u.serviceRepo.GetAll(payload.Limit, payload.Offset)
	if err != nil {
		return services, err
	}

	return services, nil
}

func (u *serviceUseCase) GetService(id string) (entity.Service, error) {
	service, err := u.serviceRepo.Get(id)
	if err != nil {
		return service, err
	}

	return service, nil
}

func (u *serviceUseCase) CreateService(service entity.Service) (entity.Service, error) {
	service.Id = common.GenerateRandomId("S")

	service, err := u.serviceRepo.Create(service)
	if err != nil {
		return service, err
	}

	return service, nil
}

func (u *serviceUseCase) UpdateService(id string, service entity.Service) (entity.Service, error) {
	service, err := u.serviceRepo.Update(id, service)
	if err != nil {
		return service, err
	}

	return service, nil
}

func (u *serviceUseCase) DeleteService(id string) error {
	err := u.serviceRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewServiceUseCase(serviceRepo repository.ServiceRepository) ServiceUseCase {
	return &serviceUseCase{
		serviceRepo: serviceRepo,
	}
}
