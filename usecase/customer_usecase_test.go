package usecase

import (
	repomock "roomate/mock/repo_mock"
	"roomate/model/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CustomerUseCaseTestSuite struct {
	suite.Suite
	crm *repomock.CustomerRepoMock
	cu  CustomerUseCase
}

func (suite *CustomerUseCaseTestSuite) SetupTest() {
	suite.crm = new(repomock.CustomerRepoMock)
	suite.cu = NewCustomerUseCase(suite.crm)
}

func TestCustomerUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerUseCaseTestSuite))
}

var dummyCustomer = entity.Customer{
	Id:          "1",
	Name:        "John",
	Email:       "john@example.com",
	Address:     "Jakarta",
	PhoneNumber: "08123456789",
}

func (suite *CustomerUseCaseTestSuite) TestCreateCustomer() {
	suite.crm.On("Create", dummyCustomer).Return(dummyCustomer, nil)
	customer, err := suite.cu.CreateCustomer(dummyCustomer)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyCustomer, customer)
}

func (suite *CustomerUseCaseTestSuite) TestGetCustomer() {
	suite.crm.On("Get", dummyCustomer.Id).Return(dummyCustomer, nil)
	customer, err := suite.cu.GetCustomer(dummyCustomer.Id)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyCustomer, customer)
}

func (suite *CustomerUseCaseTestSuite) TestGetAllCustomers() {
	suite.crm.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.Customer{dummyCustomer}, nil)
	customers, err := suite.cu.GetAllCustomers(dummyGetAllParams)
	suite.Require().NoError(err)
	suite.Require().Equal([]entity.Customer{dummyCustomer}, customers)
}

func (suite *CustomerUseCaseTestSuite) TestUpdateCustomer() {
	suite.crm.On("Update", dummyCustomer.Id, dummyCustomer).Return(dummyCustomer, nil)
	customer, err := suite.cu.UpdateCustomer(dummyCustomer.Id, dummyCustomer)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyCustomer, customer)
}

func (suite *CustomerUseCaseTestSuite) TestDeleteCustomer() {
	suite.crm.On("Delete", dummyCustomer.Id).Return(nil)
	err := suite.cu.DeleteCustomer(dummyCustomer.Id)
	suite.Require().NoError(err)
}
