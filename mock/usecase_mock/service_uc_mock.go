package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type ServiceUseCaseMock struct {
	mock.Mock
}

func (s *ServiceUseCaseMock) GetAllServices(payload dto.GetAllParams) ([]entity.Service, error) {
	args := s.Called(payload)
	return args.Get(0).([]entity.Service), args.Error(1)
}

func (s *ServiceUseCaseMock) GetService(id string) (entity.Service, error) {
	args := s.Called(id)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceUseCaseMock) CreateService(service entity.Service) (entity.Service, error) {
	args := s.Called(service)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceUseCaseMock) UpdateService(id string, service entity.Service) (entity.Service, error) {
	args := s.Called(id, service)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceUseCaseMock) DeleteService(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
