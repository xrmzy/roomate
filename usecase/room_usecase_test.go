package usecase

import (
	repomock "roomate/mock/repo_mock"
	"roomate/model/entity"
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
	dummyRoom.Id = roomId
	suite.rrm.On("Create", dummyRoom).Return(dummyRoom, nil)
	room, err := suite.ru.CreateRoom(dummyRoom)
	suite.Require().NoError(err)
	suite.Equal(dummyRoom, room)
}

func (suite *RoomUseCaseTestSuite) TestGetRoom() {
	suite.rrm.On("Get", roomId).Return(dummyRoom, nil)
	room, err := suite.ru.GetRoom(roomId)
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
	suite.rrm.On("Update", roomId, dummyRoom).Return(dummyRoom, nil)
	room, err := suite.ru.UpdateRoom(roomId, dummyRoom)
	suite.Require().NoError(err)
	suite.Equal(dummyRoom, room)
}

func (suite *RoomUseCaseTestSuite) TestUpdateStatus() {
	suite.rrm.On("UpdateStatus", roomId).Return(nil)
	err := suite.ru.UpdateStatus(roomId)
	suite.Require().NoError(err)
}

func (suite *RoomUseCaseTestSuite) TestDeleteRoom() {
	suite.rrm.On("Delete", roomId).Return(nil)
	err := suite.ru.DeleteRoom(roomId)
	suite.Require().NoError(err)
}
