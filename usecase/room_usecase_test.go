package usecase

import (
	repomock "roomate/mock/repo_mock"
	"roomate/model/entity"
	"roomate/utils/common"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RoomUseCaseTestSuite struct {
	suite.Suite
	rrm *repomock.RoomRepoMock
	ru  RoomUseCase
}

func (suite *RoomUseCaseTestSuite) SetupTest() {
	suite.rrm = new(repomock.RoomRepoMock)
	suite.ru = NewRoomUseCase(suite.rrm)
}

func TestRoomUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomUseCaseTestSuite))
}

func (suite *RoomUseCaseTestSuite) TestCreateRoom() {
	dummyRoom.Id = common.GenerateRandomId("R")
	suite.rrm.On("Create", dummyRoom).Return(dummyRoom, nil)
	room, err := suite.ru.CreateRoom(dummyRoom)
	suite.Require().NoError(err)
	suite.Equal(dummyRoom, room)
}

func (suite *RoomUseCaseTestSuite) TestGetRoom() {
	suite.rrm.On("Get", dummyRoom.Id).Return(dummyRoom, nil)
	room, err := suite.ru.GetRoom(dummyRoom.Id)
	suite.Require().NoError(err)
	suite.Equal(dummyRoom, room)
}

func (suite *RoomUseCaseTestSuite) TestGetAllRooms() {
	suite.rrm.On("GetAll", dummyGetAllParams.Limit, dummyGetAllParams.Offset).Return([]entity.Room{dummyRoom}, nil)
	rooms, err := suite.ru.GetAllRooms(dummyGetAllParams)
	suite.Require().NoError(err)
	suite.Equal([]entity.Room{dummyRoom}, rooms)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoom() {
	suite.rrm.On("Update", dummyRoom.Id, dummyRoom).Return(dummyRoom, nil)
	room, err := suite.ru.UpdateRoom(dummyRoom.Id, dummyRoom)
	suite.Require().NoError(err)
	suite.Equal(dummyRoom, room)
}

func (suite *RoomUseCaseTestSuite) TestUpdateStatus() {
	suite.rrm.On("UpdateStatus", dummyRoom.Id).Return(nil)
	err := suite.ru.UpdateStatus(dummyRoom.Id)
	suite.Require().NoError(err)
}

func (suite *RoomUseCaseTestSuite) TestDeleteRoom() {
	suite.rrm.On("Delete", dummyRoom.Id).Return(nil)
	err := suite.ru.DeleteRoom(dummyRoom.Id)
	suite.Require().NoError(err)
}
