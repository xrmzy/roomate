package repository

import (
	"database/sql"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    RoomRepository
}

// setup test
func (suite *RoomRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewRoomRepository(suite.mockDB)
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

// dummy room
var dummyRoom = entity.Room{
	Id:         "1",
	RoomNumber: "101",
	RoomType:   "Premier",
	Capacity:   2,
	Facility:   "AC, TV, Wifi",
	Price:      100000,
	Status:     "available",
	UpdatedAt:  time.Now().Truncate(time.Second),
}

func (suite *RoomRepositoryTestSuite) TestCreateRoom() {
	// design returning rows
	rows := sqlmock.NewRows([]string{"id", "room_number", "room_type", "capacity", "facility", "price", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.CreatedAt, dummyRoom.UpdatedAt, dummyRoom.IsDeleted)

	suite.sqlmock.ExpectQuery("INSERT INTO rooms").WithArgs(dummyRoom.Id, dummyRoom.RoomNumber, dummyRoom.RoomType, dummyRoom.Capacity, dummyRoom.Facility, dummyRoom.Price, dummyRoom.Status, dummyRoom.UpdatedAt).WillReturnRows(rows)

	actual, err := suite.repo.Create(dummyRoom)

	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), suite.sqlmock.ExpectationsWereMet())
	assert.Equal(suite.T(), dummyRoom, actual)
}
