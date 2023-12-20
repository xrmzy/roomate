package usecase

import (
	repomock "roomate/mock/repo_mock"
	"roomate/model/entity"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoleUseCaseTestSuite struct {
	suite.Suite
	rrm *repomock.RoleRepoMock
	ru  RoleUseCase
}

func (suite *RoleUseCaseTestSuite) SetupTest() {
	suite.rrm = new(repomock.RoleRepoMock)
	suite.ru = NewRoleUseCase(suite.rrm)
}

func TestRoleUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoleUseCaseTestSuite))
}

var dummyRole = entity.Role{
	Id:       "1",
	RoleName: "admin",
}

func (suite *RoleUseCaseTestSuite) TestCreateRole() {
	suite.rrm.On("Create", dummyRole).Return(dummyRole, nil)
	role, err := suite.ru.CreateRole(dummyRole)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyRole, role)
}

func (suite *RoleUseCaseTestSuite) TestGetRole() {
	var id string = "1"
	suite.rrm.On("Get", id).Return(dummyRole, nil)
	role, err := suite.ru.GetRole(id)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyRole, role)
}

func (suite *RoleUseCaseTestSuite) TestGetAllRoles() {
	suite.rrm.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.Role{dummyRole}, nil)
	roles, err := suite.ru.GetAllRoles(dummyGetAllParams)
	suite.Require().NoError(err)
	suite.Require().Equal([]entity.Role{dummyRole}, roles)
}

func (suite *RoleUseCaseTestSuite) TestUpdateRole() {
	var id string = "1"
	suite.rrm.On("Update", id, dummyRole).Return(dummyRole, nil)
	role, err := suite.ru.UpdateRole(id, dummyRole)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyRole, role)
}

func (suite *RoleUseCaseTestSuite) TestDeleteRole() {
	var id string = "1"
	suite.rrm.On("Delete", id).Return(nil)
	err := suite.ru.DeleteRole(id)
	suite.Require().NoError(err)
}
