package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"roomate/delivery/middleware"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	uum    *usecasemock.MockUserUseCase
	engine *gin.Engine
	amm    *middleware.AuthMiddleware
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.uum = new(usecasemock.MockUserUseCase)
	suite.engine = gin.Default()
	suite.amm = new(middleware.AuthMiddleware)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

var dummyUser = entity.User{
	Id:        "1",
	Name:      "Dimas",
	Email:     "dimas@gmail.com",
	Password:  "Wadaw",
	RoleId:    "1",
	RoleName:  "Admin",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	IsDeleted: false,
}

func (suite *UserControllerTestSuite) TestCreateUser_Success() {
	suite.uum.On("CreateUser", mock.AnythingOfType("entity.User")).Return(dummyUser, nil)

	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	mockPayloadJSON, err := json.Marshal(dummyUser)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusCreated, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestCreateUser_Fail() {
	// Mock data dan ekspetasi untuk kasus kegagalan
	suite.uum.On("CreateUser", mock.AnythingOfType("entity.User")).Return(entity.User{}, errors.New("failed to create User"))

	// Setup router dan handler
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	// Membuat request untuk endpoint create user
	mockPayloadJSON, err := json.Marshal(dummyUser)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat recorder untuk menyimpan respons
	w := httptest.NewRecorder()

	// Menangani request dengan router
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

//func

func (suite *UserControllerTestSuite) TestGetUserHandler_Success() {
	suite.uum.On("GetUser", "1").Return(dummyUser, nil)

	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	suite.NoError(err)

	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

// func
func (suite *UserControllerTestSuite) TestGetAllUserHandler_Success() {
	suite.uum.On("GetAllUsers", mock.AnythingOfType("dto.GetAllParams")).Return([]entity.User{
		{Id: "S001", Name: "Dimas", Email: "dimas@gmail,com", Password: "asasas", RoleId: "1", RoleName: "Admin"},
		{Id: "S002", Name: "Saha", Email: "saha@gmail.com", Password: "adadadsaf", RoleId: "2", RoleName: "Employee"},
	}, nil)
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"offset": 0, "limit": 10}`)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

// func
func (suite *UserControllerTestSuite) TestUpdateUserHandler_Success() {
	suite.uum.On("UpdateUser", mock.AnythingOfType("string"), mock.AnythingOfType("entity.User")).Return(dummyUser, nil)

	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	mockPayloadJSON := []byte(`{"name": "Sala", "email": "saalap@gmail.com"}`)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/1", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)
	fmt.Println("Response Body:", w.Body.String())

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

// func
func (suite *UserControllerTestSuite) TestUpdateUserHandler_Failure() {
	// EKSPETASI
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1/"))
	userController.Route()

	// Membuat request dengan payload yang tidak valid
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/1", bytes.NewBufferString("invalid json"))
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusBadRequest, w.Code)

	// Tidak ada pemanggilan metode pada objek mock (payload tidak valid)
	suite.uum.AssertExpectations(suite.T())
}

// func
func (suite *UserControllerTestSuite) TestDeleteUserHandler_Success() {
	// Service ID yang akan dihapus
	userId := "1"

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.uum.On("DeleteUser", userId).Return(nil)

	// Membuat ServiceController instance
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/"+userId, nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestDeleteUserHandler_Failure() {
	suite.uum.On("DeleteUser", mock.AnythingOfType("string")).Return(errors.New("failed to delete user"))

	// Membuat RoleController instance
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestUpdatePassword_Success() {
	// Konfigurasi mock untuk metode UpdateUser
	suite.uum.On("UpdateUser", mock.AnythingOfType("string"), mock.AnythingOfType("entity.User")).Return(dummyUser, nil)

	// Membuat UserController instance
	userController := NewUserController(suite.uum, suite.engine.Group("/api/v1"))
	userController.Route()

	// Persiapkan data payload JSON
	updatePasswordData := map[string]string{
		"id":       "1",
		"password": "NewPassword",
	}
	mockPayloadJSON, err := json.Marshal(updatePasswordData)
	suite.NoError(err)

	// Buat HTTP request
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/update-password:1", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Buat response recorder
	w := httptest.NewRecorder()

	// Jalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.uum.AssertExpectations(suite.T())
}
