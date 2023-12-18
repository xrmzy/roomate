package usecase

import (
	repomock "roomate/mock/repo_mock"
	"roomate/model/entity"
	"roomate/utils/common"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceUseCaseTestSuite struct {
	suite.Suite
	srm *repomock.ServiceRepoMock
	su  ServiceUseCase
}

func (suite *ServiceUseCaseTestSuite) SetupTest() {
	suite.srm = new(repomock.ServiceRepoMock)
	suite.su = NewServiceUseCase(suite.srm)
}

func TestServiceUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceUseCaseTestSuite))
}

func (suite *ServiceUseCaseTestSuite) TestCreateService() {
	dummyService.Id = common.GenerateRandomId("R")
	suite.srm.On("Create", dummyService).Return(dummyService, nil)
	service, err := suite.su.CreateService(dummyService)
	suite.Require().NoError(err)
	suite.Equal(dummyService, service)
}

func (suite *ServiceUseCaseTestSuite) TestGetService() {
	suite.srm.On("Get", dummyService.Id).Return(dummyService, nil)
	service, err := suite.su.GetService(dummyService.Id)
	suite.Require().NoError(err)
	suite.Equal(dummyService, service)
}

func (suite *ServiceUseCaseTestSuite) TestGetAllService() {
	suite.srm.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.Service{dummyService}, nil)
	service, err := suite.su.GetAllServices(dummyGetAllParams)
	suite.Require().NoError(err)
	suite.Equal([]entity.Service{dummyService}, service)
}

func (suite *ServiceUseCaseTestSuite) TestUpdateService() {
	suite.srm.On("Update", dummyService.Id, dummyService).Return(dummyService, nil)
	service, err := suite.su.UpdateService(dummyService.Id, dummyService)
	suite.Require().NoError(err)
	suite.Equal(dummyService, service)
}

func (suite *ServiceUseCaseTestSuite) TestDeleteService() {
	suite.srm.On("Delete", dummyService.Id).Return(nil)
	err := suite.su.DeleteService(dummyService.Id)
	suite.Require().NoError(err)
}
