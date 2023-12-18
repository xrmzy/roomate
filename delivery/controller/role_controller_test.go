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

// RoleControllerTestSuite adalah suite pengujian untuk RoleController
type RoleControllerTestSuite struct {
	suite.Suite
	rum    *usecasemock.RoleUseCaseMock
	engine *gin.Engine
}

// SetupTest digunakan untuk melakukan inisialisasi sebelum setiap pengujian
func (suite *RoleControllerTestSuite) SetupTest() {
	suite.rum = new(usecasemock.RoleUseCaseMock)
	suite.engine = gin.Default()
}

func TestRoleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoleControllerTestSuite))
}

var dummyRole = entity.Role{
	Id:        1,
	RoleName:  "Admin",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	IsDeleted: false,
}

// TestCreateHandler_Success menguji skenario ketika pembuatan role berhasil
func (suite *RoleControllerTestSuite) TestCreateHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.rum.On("CreateRole", mock.AnythingOfType("entity.Role")).Return(dummyRole, nil)

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	mockPayloadJSON, err := json.Marshal(dummyRole)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/roles", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusCreated, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

// TestCreateHandler_Fail menguji skenario ketika pembuatan role gagal
func (suite *RoleControllerTestSuite) TestCreateHandler_Fail() {
	// Menyusun ekspetasi panggilan fungsi di UseCase untuk mengembalikan kesalahan
	suite.rum.On("CreateRole", mock.AnythingOfType("entity.Role")).Return(entity.Role{}, errors.New("failed to create role"))

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	mockPayloadJSON, err := json.Marshal(dummyRole)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/roles", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

func (suite *RoleControllerTestSuite) TestGetHandler_Success() {

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.rum.On("GetRole", mock.AnythingOfType("string")).Return(dummyRole, nil)

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/roles/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

func (suite *RoleControllerTestSuite) TestGetAllHandler_Success() {

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.rum.On("GetAllRoles", mock.AnythingOfType("dto.GetAllParams")).Return([]entity.Role{
		{Id: 1, RoleName: "Admin"},
		{Id: 2, RoleName: "Employee"},
	}, nil)

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"offset": 0, "limit": 10}`)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/roles", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

func (suite *RoleControllerTestSuite) TestUpdateHandler_Success() {

	// Parameter ID yang ingin diupdate
	roleId := "1"

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.rum.On("UpdateRole", roleId, mock.AnythingOfType("entity.Role")).Return(dummyRole, nil)

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"roleName": "UpdatedRole"}`)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/roles/"+roleId, bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

// func
func (suite *RoleControllerTestSuite) TestUpdateRoleHandler_Failure() {
	// EKSPETASI
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1/"))
	roleController.Route()

	// Membuat request dengan payload yang tidak valid
	req, err := http.NewRequest(http.MethodPut, "/api/v1/roles/1", bytes.NewBufferString("invalid json"))
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusBadRequest, w.Code)

	// Tidak ada pemanggilan metode pada objek mock (payload tidak valid)
	suite.rum.AssertExpectations(suite.T())
}

func (suite *RoleControllerTestSuite) TestDeleteHandler_Success() {

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.rum.On("DeleteRole", mock.AnythingOfType("string")).Return(nil)

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/roles/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}

func (suite *RoleControllerTestSuite) TestDeleteHandler_Fail() {
	// Menyusun ekspetasi panggilan fungsi di UseCase untuk mengembalikan kesalahan
	suite.rum.On("DeleteRole", mock.AnythingOfType("string")).Return(errors.New("failed to delete role"))

	// Membuat RoleController instance
	roleController := NewRoleController(suite.rum, suite.engine.Group("/api/v1"))
	roleController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/roles/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.rum.AssertExpectations(suite.T())
}
