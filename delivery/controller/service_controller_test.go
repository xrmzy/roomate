package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

type ServiceControllerTestSuite struct {
	suite.Suite
	sum    *usecasemock.ServiceUseCaseMock
	engine *gin.Engine
}

func (suite *ServiceControllerTestSuite) SetupTest() {
	suite.sum = new(usecasemock.ServiceUseCaseMock)
	suite.engine = gin.Default()
}

func TestServiceControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceControllerTestSuite))
}

var dummyService = entity.Service{
	Id:        "S091",
	Name:      "Proyektor",
	Price:     20000,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	IsDeleted: false,
}

func (suite *ServiceControllerTestSuite) TestCreateService_Success() {
	suite.sum.On("CreateService", mock.AnythingOfType("entity.Service")).Return(dummyService, nil)
	serviceControllers := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceControllers.Route()

	mockPayloadJSON, err := json.Marshal(dummyService)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/services", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusCreated, w.Code)
	fmt.Println("Response Body:", w.Body.String())

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}

func (suite *ServiceControllerTestSuite) TestCreateServiceHandler_Fail() {
	// Menyusun ekspetasi panggilan fungsi di UseCase untuk mengembalikan kesalahan
	suite.sum.On("CreateService", mock.AnythingOfType("entity.Service")).Return(entity.Service{}, errors.New("failed to create Service"))

	// Membuat RoleController instance
	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceController.Route()

	// Membuat request
	mockPayloadJSON, err := json.Marshal(dummyService)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/services", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}

// func (suite *ServiceControllerTestSuite) a()                             {}
func (suite *ServiceControllerTestSuite) TestGetServiceHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.sum.On("GetService", "S091").Return(dummyService, nil)

	// Membuat ServiceController instance
	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1/"))
	serviceController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/services/S091", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)
	fmt.Println("Response Body:", w.Body.String())

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}

func (suite *ServiceControllerTestSuite) TestGetAllHandler_Success() {
	// Ekpetasi
	suite.sum.On("GetAllServices", mock.AnythingOfType("dto.GetAllParams")).Return([]entity.Service{
		{Id: "S001", Name: "Proyektor", Price: 20000},
		{Id: "S002", Name: "Seprai", Price: 15000},
	}, nil)

	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"offset": 0, "limit": 10}`)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/services", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())

}

func (suite *ServiceControllerTestSuite) TestUpdateHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.sum.On("UpdateService", mock.AnythingOfType("string"), mock.AnythingOfType("entity.Service")).
		Return(dummyService, nil)

	// Membuat ServiceController instance
	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"name": "Handuk", "price": 25000}`)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/services/S091", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}

func (suite *ServiceControllerTestSuite) TestDeleteHandler_Success() {
	// Service ID yang akan dihapus
	serviceId := "S091"

	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.sum.On("DeleteService", serviceId).Return(nil)

	// Membuat ServiceController instance
	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/services/"+serviceId, nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}

// fi=unc
func (suite *ServiceControllerTestSuite) TestDeleteServiceHandler_Failure() {
	suite.sum.On("DeleteService", mock.AnythingOfType("string")).Return(errors.New("failed to delete service"))

	// Membuat RoleController instance
	serviceController := NewServiceController(suite.sum, suite.engine.Group("/api/v1"))
	serviceController.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/services/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.sum.AssertExpectations(suite.T())
}
