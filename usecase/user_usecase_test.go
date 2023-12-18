package usecase

import (
	"errors"
	commonmock "roomate/mock/common_mock"
	repomock "roomate/mock/repo_mock"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	userRepoMock *repomock.UserRepoMock
	roleUcMock   *usecasemock.RoleUseCaseMock
	pHashMock    *commonmock.PasswordHashCommonMock
	userUc       UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.userRepoMock = new(repomock.UserRepoMock)
	suite.roleUcMock = new(usecasemock.RoleUseCaseMock)
	suite.pHashMock = new(commonmock.PasswordHashCommonMock)
	suite.userUc = NewUserUseCase(suite.userRepoMock, suite.roleUcMock)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) TestGetUser() {
	// get user fail
	suite.userRepoMock.On("Get", dummyUser.Id).Return(entity.User{}, errors.New("user not found")).Once()
	_, err := suite.userUc.GetUser(dummyUser.Id)
	suite.Require().Error(err)

	suite.userRepoMock.On("Get", dummyUser.Id).Return(dummyUser, nil)
	room, err := suite.userUc.GetUser(dummyUser.Id)
	suite.Require().NoError(err)
	suite.Equal(dummyUser, room)
}

func (suite *UserUseCaseTestSuite) TestGetAllUser() {
	// get all users err
	suite.userRepoMock.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.User{}, errors.New("user not found")).Once()
	_, err := suite.userUc.GetAllUsers(dummyGetAllParams)
	suite.Require().Error(err)

	suite.userRepoMock.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.User{dummyUser}, nil)
	users, err := suite.userUc.GetAllUsers(dummyGetAllParams)
	suite.Require().NoError(err)
	suite.Equal([]entity.User{dummyUser}, users)
}

func (suite *UserUseCaseTestSuite) TestUpdateUser() {
	// update user err
	suite.userRepoMock.On("Update", dummyUser.Id, dummyUser).Return(entity.User{}, errors.New("update user error")).Once()
	_, err := suite.userUc.UpdateUser(dummyUser.Id, dummyUser)
	suite.Require().Error(err)

	suite.userRepoMock.On("Update", dummyUser.Id, dummyUser).Return(dummyUser, nil)
	user, err := suite.userUc.UpdateUser(dummyUser.Id, dummyUser)
	suite.Require().NoError(err)
	suite.Equal(dummyUser, user)
}

func (suite *UserUseCaseTestSuite) TestDeleteUser() {
	// delete user err
	suite.userRepoMock.On("Delete", dummyUser.Id).Return(errors.New("delete user error")).Once()
	err := suite.userUc.DeleteUser(dummyUser.Id)
	suite.Require().Error(err)

	suite.userRepoMock.On("Delete", dummyUser.Id).Return(nil)
	err = suite.userUc.DeleteUser(dummyUser.Id)
	suite.Require().NoError(err)
}

func (suite *UserUseCaseTestSuite) TestGetByEmailPassword() {
	// dummyUser
	var dummyUser = entity.User{
		Id:        "1",
		Name:      "John",
		Email:     "johndoe@me.com",
		Password:  "$2a$10$uezZkgN6CtY/UWll7MLT8Os0y2m87GnI6druwXtnU.cNnO7.Y.LKW",
		RoleId:    "1",
		RoleName:  "admin",
		UpdatedAt: time.Now().Truncate(time.Second),
	}
	var password = "167916"

	suite.userRepoMock.On("GetByEmail", dummyUser.Email).Return(dummyUser, nil).Once()

	suite.pHashMock.On("ComparePasswordHash", dummyUser.Password, password).Return(nil).Once()

	userr, err := suite.userUc.GetByEmailPassword(dummyUser.Email, password)
	dummyUser.Password = ""
	suite.Require().NoError(err)
	suite.Equal(dummyUser, userr)
}

// func (suite *UserUseCaseTestSuite) TestCreateUser() {
// 	suite.userRepoMock.On("GetByEmail", dummyUser.Email).Return(entity.User{}, nil).Once()

// 	suite.roleUcMock.On("GetRole", dummyUser.RoleId).Return(dummyRole, nil).Once()

// 	// dummyUser.Password, _ = common.GeneratePasswordHash(dummyUser.Password)
// 	suite.userRepoMock.On("Create", dummyUser).Return(dummyUser, nil)

// 	user, err := suite.userUc.CreateUser(dummyUser)
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(dummyUser, user)
// }