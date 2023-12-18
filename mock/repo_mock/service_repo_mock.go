package repomock

import (
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type ServiceRepoMock struct {
	mock.Mock
}

func (s *ServiceRepoMock) Get(id string) (entity.Service, error) {
	args := s.Called(id)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceRepoMock) GetAll(limit, offset int) ([]entity.Service, error) {
	args := s.Called(limit, offset)
	return args.Get(0).([]entity.Service), args.Error(1)
}

func (s *ServiceRepoMock) Create(service entity.Service) (entity.Service, error) {
	args := s.Called(service)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceRepoMock) Update(id string, service entity.Service) (entity.Service, error) {
	args := s.Called(id, service)
	return args.Get(0).(entity.Service), args.Error(1)
}

func (s *ServiceRepoMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
