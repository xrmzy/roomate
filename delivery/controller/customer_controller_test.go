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
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CustomerControllerTestSuite struct {
	suite.Suite
	cum    *usecasemock.CustomerUseCaseMock
	engine *gin.Engine
}

func (suite *CustomerControllerTestSuite) SetupTest() {
	suite.cum = new(usecasemock.CustomerUseCaseMock)
	suite.engine = gin.Default()
}

func TestCustomerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerControllerTestSuite))
}

var dummyCustomer = entity.Customer{
	Id:          "1",
	Name:        "Dimas",
	Email:       "dimas@gmail.com",
	Address:     "Serang",
	PhoneNumber: "091212",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	IsDeleted:   false,
}

func buildJSONPayload(customer entity.Customer) string {
	payload := fmt.Sprintf(`{"name": "%s", "email": "%s", "address": "%s", "phoneNumber": "%s"}`,
		customer.Name, customer.Email, customer.Address, customer.PhoneNumber)
	return payload
}

func (suite *CustomerControllerTestSuite) TestCreateCustomer_Testing() {
	suite.cum.On("CreateCustomer", mock.AnythingOfType("entity.Customer")).Return(dummyCustomer, nil)
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1"))
	customerController.Route()

	mockPayloadJSON, err := json.Marshal(dummyRoom)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusCreated, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestCreateHandler_Fail() {
	// Menyusun ekspetasi panggilan fungsi di UseCase untuk mengembalikan kesalahan
	suite.cum.On("CreateCustomer", mock.AnythingOfType("entity.Customer")).Return(entity.Customer{}, errors.New("failed to create role"))

	// Membuat CustomerController instance
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1"))
	customerController.Route()

	// Membuat request
	mockPayloadJSON, err := json.Marshal(dummyCustomer)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (gagal)
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

// func
func (suite *CustomerControllerTestSuite) TestGetCustomersHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.cum.On("GetCustomer", "1").Return(dummyCustomer, nil)

	// Membuat CustomerController instance
	customerControllers := NewCustomerController(suite.cum, suite.engine.Group("/api/v1/"))
	customerControllers.Route()

	// Membuat request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customers/1", nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

//dd

func (suite *CustomerControllerTestSuite) TestUpdateCustomerHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.cum.On("UpdateCustomer", "1", mock.AnythingOfType("entity.Customer")).Return(dummyCustomer, nil)

	// Membuat CustomerController instance
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1/"))
	customerController.Route()

	// Membuat request dengan data yang akan diupdate
	jsonStr := `{"name": "Dimas", "email": "dimas_sasasa@gmail.com", "address": "Bandung", "phoneNumber": "123456789"}`
	req, err := http.NewRequest(http.MethodPut, "/api/v1/customers/1", strings.NewReader(jsonStr))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestDeleteCustomerHandler_Success() {
	// Menyusun ekspetasi panggilan fungsi di UseCase
	suite.cum.On("DeleteCustomer", "1").Return(nil)

	// Membuat CustomerController instance
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1/"))
	customerController.Route()

	// Membuat request untuk delete
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/customers/%s", dummyCustomer.Id), nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

func (suite *CustomerControllerTestSuite) TestUpdateCustomerHandler_Failure() {
	// Menyusun ekspetasi panggilan fungsi di UseCase yang mengembalikan error
	expectedError := errors.New("Update customer error")
	suite.cum.On("UpdateCustomer", "1", mock.AnythingOfType("entity.Customer")).Return(entity.Customer{}, expectedError)

	// Membuat CustomerController instance
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1/"))
	customerController.Route()

	// Membuat request dengan data yang akan diupdate
	jsonStr := buildJSONPayload(dummyCustomer)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/customers/%s", dummyCustomer.Id), strings.NewReader(jsonStr))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (seharusnya mengembalikan status error)
	suite.Equal(http.StatusInternalServerError, w.Code) // Ganti status code dengan yang diharapkan

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

// TestDeleteCustomerHandler_Failure melakukan unit testing pada kasus delete yang gagal
func (suite *CustomerControllerTestSuite) TestDeleteCustomerHandler_Failure() {
	// Menyusun ekspetasi panggilan fungsi di UseCase yang mengembalikan error
	expectedError := errors.New("Delete customer error")
	suite.cum.On("DeleteCustomer", "1").Return(expectedError)

	// Membuat CustomerController instance
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1/"))
	customerController.Route()

	// Membuat request untuk delete
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/customers/%s", dummyCustomer.Id), nil)
	suite.NoError(err)

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan (seharusnya mengembalikan status error)
	suite.Equal(http.StatusInternalServerError, w.Code) // Ganti status code dengan yang diharapkan

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}

//a

func (suite *CustomerControllerTestSuite) TestGetAllCustomerHandler_Success() {
	suite.cum.On("GetAllCustomers", mock.AnythingOfType("dto.GetAllParams")).Return([]entity.Customer{
		{Id: "S001", Name: "Dimas", Email: "dimas@gmail,com", Address: "Serang", PhoneNumber: "0812912"},
		{Id: "S002", Name: "Saha", Email: "saha@gmail.com", Address: "Bandung", PhoneNumber: "00099"},
	}, nil)
	customerController := NewCustomerController(suite.cum, suite.engine.Group("/api/v1"))
	customerController.Route()

	// Membuat request
	mockPayloadJSON := []byte(`{"offset": 0, "limit": 10}`)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customers", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Membuat response recorder
	w := httptest.NewRecorder()

	// Menjalankan request ke handler
	suite.engine.ServeHTTP(w, req)

	// Memeriksa bahwa handler memberikan respons yang diharapkan
	suite.Equal(http.StatusOK, w.Code)

	// Menyelesaikan ekspetasi UseCase
	suite.cum.AssertExpectations(suite.T())
}
