package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"roomate/model/entity"
	"roomate/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock

	repo RoomRepository
}

func (suite *RoomRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewRoomRepository(suite.mockDB)
}

func TestRoomRepoSuite(t *testing.T) {
	suite.Run(t, new(RoomRepoTestSuite))
}

var dummyRoom = entity.Room{
	Id:         "R12345",
	RoomNumber: "45",
	RoomType:   "President",
	Capacity:   1,
	Facility:   "Pool, Meeting Room, Guest Room",
	Price:      1750000,
	Status:     "available",
	CreatedAt:  time.Now().Truncate(time.Second),
	UpdatedAt:  time.Now().Truncate(time.Second),
	IsDeleted:  false,
}

func (suite *RoomRepoTestSuite) TestRepository_GetRoom() {
	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetRoom)).WithArgs(dummyRoom.Id).WillReturnError(errors.New("Get by id failed"))

	_, err := suite.repo.Get(dummyRoom.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Get by id failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "room_number", "room_type", "capacity", "facility", "price", "status", "created_at", "updated_at"}).AddRow(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.CreatedAt, dummyRoom.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetRoom)).WithArgs(dummyRoom.Id).WillReturnRows(rows)

	actual, err := suite.repo.Get(dummyRoom.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRoom.Id, actual.Id)
}

func (suite *RoomRepoTestSuite) TestRepository_CreateRoom() {

	// test fail
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateRoom)).
		WithArgs(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.UpdatedAt).
		WillReturnError(errors.New("Create room failed"))

	_, err := suite.repo.Create(dummyRoom)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Create room failed", err.Error())

	rows := sqlmock.NewRows([]string{"id", "room_number", "room_type", "capacity", "facility", "price", "status", "created_at", "updated_at"}).
		AddRow(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.CreatedAt, dummyRoom.UpdatedAt)

	// test success
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateRoom)).
		WithArgs(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.UpdatedAt).
		WillReturnRows(rows)

	actual, err := suite.repo.Create(dummyRoom)
	fmt.Println("generate id actual" + actual.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRoom.Id, actual.Id)
}

func (suite *RoomRepoTestSuite) TestRepository_GetAllRoom() {

	// test failed

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllRooms)).WithArgs(1, 0).WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{"id", "room_number", "room_type", "capacity", "facility", "price", "status", "created_at", "updated_at"}).
		AddRow(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.CreatedAt, dummyRoom.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllRooms)).WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRoom.Id, actual[0].Id)
	assert.Equal(suite.T(), dummyRoom.RoomNumber, actual[0].RoomNumber)

}

func (suite *RoomRepoTestSuite) TestRepository_UpdateRoom() {

	updatedRoom := dummyRoom
	updatedRoom.RoomNumber = "45"
	updatedRoom.RoomType = "Single room"
	updatedRoom.Capacity = 1
	updatedRoom.Facility = "desk, chair, TV, and a small bathroom "
	updatedRoom.Price = 700000
	updatedRoom.Status = "Booked"

	// mock error from QueryRow
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateRoom)).WithArgs(dummyRoom.Id, updatedRoom.RoomNumber, updatedRoom.RoomType, updatedRoom.Capacity, updatedRoom.Facility, updatedRoom.Price, updatedRoom.Status, updatedRoom.UpdatedAt).WillReturnError(errors.New("update failed"))

	_, err := suite.repo.Update(dummyRoom.Id, updatedRoom)

	// assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update failed", err.Error())

	// test success
	rows := sqlmock.NewRows([]string{"id", "room_number", "room_type", "capacity", "facility", "price", "status", "created_at", "updated_at"}).
		AddRow(dummyRoom.Id, updatedRoom.RoomNumber, updatedRoom.RoomType, updatedRoom.Capacity, updatedRoom.Facility, updatedRoom.Price, updatedRoom.Status, updatedRoom.CreatedAt, updatedRoom.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateRoom)).WithArgs(dummyRoom.Id, updatedRoom.RoomNumber, updatedRoom.RoomType, updatedRoom.Capacity, updatedRoom.Facility, updatedRoom.Price, updatedRoom.Status, updatedRoom.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyRoom.Id, updatedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyRoom.Id, actual.Id)
	assert.Equal(suite.T(), updatedRoom.RoomNumber, actual.RoomNumber)
	assert.Equal(suite.T(), updatedRoom.RoomType, actual.RoomType)
	assert.Equal(suite.T(), updatedRoom.Status, actual.Status)

}

func (suite *RoomRepoTestSuite) TestRepository_DeleteRoom() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteRoom)).WithArgs(dummyRoom.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyRoom.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteRoom)).WithArgs(dummyRoom.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyRoom.Id)

	assert.NoError(suite.T(), err)
}
