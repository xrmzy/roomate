package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	jwtmock "roomate/mock/jwt_mock"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	aum    *usecasemock.MockAuthUseCase
	engine *gin.Engine
	jmm    *jwtmock.MockJwtToken
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.aum = new(usecasemock.MockAuthUseCase)
	suite.jmm = new(jwtmock.MockJwtToken)
	suite.engine = gin.Default()
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

var dummyLogin = dto.AuthRequestDto{
	Email:    "dimas@gmail.com",
	Password: "wadaw",
}

var dummyToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI5MTE3OTgsImlhdCI6MTcwMjgyNTM5OCwidXNlcklkIjoiZmI5ZWJkYTUtZjE5MS00MzA0LTljZWYtODdiOWQyOWM3YjY4Iiwicm9sZSI6IkFkbWluIn0.qKOt7OwRAiJzcd037KwuDgE4qhP2rxkO40TlXKfIhmE"

var dummyRefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI3OTExMzcsImlhdCI6MTcwMjcwNDczNywidXNlcklkIjoiZmI5ZWJkYTUtZjE5MS00MzA0LTljZWYtODdiOWQyOWM3YjY4Iiwicm9sZSI6IkFkbWluIn0.jg71NzO_UCb2MQXE_R12AnM_gwmEP9n7VkktL1Y2hxs"

func (suite *AuthControllerTestSuite) TestLoginSuccess() {
	// Configure the mock for the Login method
	suite.aum.On("Login", mock.AnythingOfType("dto.AuthRequestDto")).
		Return(dto.AuthResponseDto{Token: dummyToken}, nil)

	// Create an AuthController instance
	authController := NewAuthController(suite.aum, suite.engine.Group("/api/v1/"), suite.jmm)
	authController.Route()

	// Prepare JSON payload
	mockPayloadJSON, err := json.Marshal(dummyLogin)
	suite.NoError(err)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/login", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Execute the request
	suite.engine.ServeHTTP(w, req)

	// Check that the handler provides the expected response
	suite.Equal(http.StatusOK, w.Code)

	// Complete the mock expectations
	suite.aum.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestLoginFail() {
	// Configure the mock for the Login method to return an error
	suite.aum.On("Login", mock.AnythingOfType("dto.AuthRequestDto")).
		Return(dto.AuthResponseDto{}, errors.New("authentication failed"))

	// Create an AuthController instance
	authController := NewAuthController(suite.aum, suite.engine.Group("/api/v1/"), suite.jmm)
	authController.Route()

	// Prepare JSON payload
	mockPayloadJSON, err := json.Marshal(dummyLogin)
	suite.NoError(err)

	// Create an HTTP request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/login", bytes.NewBuffer(mockPayloadJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Execute the request
	suite.engine.ServeHTTP(w, req)

	// Check that the handler provides the expected response
	suite.Equal(http.StatusInternalServerError, w.Code)

	// Complete the mock expectations
	suite.aum.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRefreshTokenSuccess() {
	// Configure the mock for the RefreshToken method
	suite.jmm.On("RefreshToken", dummyRefreshToken).
		Return(dto.AuthResponseDto{Token: dummyToken}, nil)

	// Create an AuthController instance
	authController := NewAuthController(suite.aum, suite.engine.Group("/api/v1/"), suite.jmm)
	authController.Route()

	// Prepare HTTP request for refresh token
	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/refresh-token", nil)
	suite.NoError(err)
	req.Header.Set("Authorization", "Bearer "+dummyRefreshToken)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Execute the request
	suite.engine.ServeHTTP(w, req)

	// Check that the handler provides the expected response
	suite.Equal(http.StatusOK, w.Code)

	// Complete the mock expectations
	suite.jmm.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRefreshTokenFail() {
	// Configure the mock for the RefreshToken method to return an error
	suite.jmm.On("RefreshToken", dummyRefreshToken).
		Return(dto.AuthResponseDto{}, errors.New("refresh token failed")) // someError harus didefinisikan sesuai dengan kasus Anda

	// Create an AuthController instance
	authController := NewAuthController(suite.aum, suite.engine.Group("/api/v1/"), suite.jmm)
	authController.Route()

	// Prepare HTTP request for refresh token
	req, err := http.NewRequest(http.MethodGet, "/api/v1/auth/refresh-token", nil)
	suite.NoError(err)
	req.Header.Set("Authorization", "Bearer "+dummyRefreshToken)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Execute the request
	suite.engine.ServeHTTP(w, req)

	// Check that the handler provides the expected response
	suite.Equal(http.StatusUnauthorized, w.Code)
}
