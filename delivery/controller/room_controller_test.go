package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RoomControllerTestSuite struct {
	suite.Suite
	roum   *usecasemock.RoomUseCaseMock
	engine *gin.Engine
}

func (suite *RoomControllerTestSuite) SetupTest() {
	suite.roum = new(usecasemock.RoomUseCaseMock)
	suite.engine = gin.Default()
}

func TestRoomControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoomControllerTestSuite))
}

var dummyRoom = entity.Room{
	Id:         "R091",
	RoomNumber: "R10",
	RoomType:   "Super",
	Capacity:   2,
	Facility:   "AC",
	Price:      120000,
	Status:     "booking",
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
	IsDeleted:  false,
}

func (suite *RoomControllerTestSuite) TestCreateRoom_Testing() {
	suite.roum.On("CreateRoom", mock.AnythingOfType("entity.Room")).Return(dummyRoom, nil)
	roomceControllers := NewRoomController(suite.roum, suite.engine.Group("/api/v1"))
	roomceControllers.Route()

	mockPayloadJSON, err := json.Marshal(dummyRoom)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/rooms", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusCreated, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestCreateRoomsHandler_Fail() {
	// Menyusun ekspetasi panggilan fungsi di UseCase untuk mengembalikan kesalahan
	suite.roum.On("CreateRoom", mock.AnythingOfType("entity.Room")).Return(entity.Room{}, errors.New("failed to create Service"))

	// Membuat romeController instance
	roomController := NewRoomController(suite.roum, suite.engine.Group("/api/v1"))
	roomController.Route()

	// Membuat request
	mockPayloadJSON, err := json.Marshal(dummyService)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/rooms", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())
}

// func
func (suite *RoomControllerTestSuite) TestGetRoomsHandler_Success() {
	// Ekspetasi
	suite.roum.On("GetRoom", "R091").Return(dummyRoom, nil)

	// Membuat RoomsController instance
	roomceControllers := NewRoomController(suite.roum, suite.engine.Group("/api/v1/"))
	roomceControllers.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/rooms/R091", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestGetAllHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.roum.On("GetAllRooms", mock.AnythingOfType("dto.GetAllParams")).Return([]entity.Room{
		{Id: "R091", RoomNumber: "R10", RoomType: "SUPER", Capacity: 2, Facility: "AC", Price: 120000, Status: "booking"},
		{Id: "R092", RoomNumber: "R12", RoomType: "eXCLUSIV", Capacity: 4, Facility: "AC", Price: 220000, Status: "AVAILABLE"},
	}, nil)

	romeController := NewRoomController(suite.roum, suite.engine.Group("/api/v1"))
	romeController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"offset": 0, "limit": 10}`)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/rooms", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())

}

func (suite *RoomControllerTestSuite) TestUpdateHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.roum.On("UpdateRoom", "R091", mock.AnythingOfType("entity.Room")).Return(dummyRoom, nil)

	// Membuat RoomController instance
	roomController := NewRoomController(suite.roum, suite.engine.Group("/api/v1/"))
	roomController.Route()

	// Membuat request
	payload := entity.Room{RoomNumber: "R11", RoomType: "DELUXE", Capacity: 3, Facility: "TV", Price: 150000, Status: "AVAILABLE"}
	payloadJSON, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/rooms/R091", bytes.NewBuffer(payloadJSON))
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestUpdateHandler_BadRequest() {
	// Membuat RoomController instance
	roomController := NewRoomController(suite.roum, suite.engine.Group("/api/v1/"))
	roomController.Route()

	// Membuat request dengan payload yang tidak valid
	req, err := http.NewRequest(http.MethodPut, "/api/v1/rooms/R091", bytes.NewBufferString("invalid json"))
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusBadRequest, w.Code)

	// Tidak ada pemanggilan metode pada objek mock (payload tidak valid)
	suite.roum.AssertExpectations(suite.T())
}

func (suite *RoomControllerTestSuite) TestDeleteHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.roum.On("DeleteRoom", "R091").Return(nil)

	// Membuat RoomController instance
	roomController := NewRoomController(suite.roum, suite.engine.Group("/api/v1/"))
	roomController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/rooms/R091", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.roum.AssertExpectations(suite.T())
}
