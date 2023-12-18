package usecase_test

import (
	"errors"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	uum *usecasemock.MockUserUseCase
	jum *usecasemock.MockJwtToken
}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.uum = &usecasemock.MockUserUseCase{}
	suite.jum = &usecasemock.MockJwtToken{}
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	mockUser := entity.User{Id: "1", Email: "test@example.com", Password: "password"}
	mockToken := dto.AuthResponseDto{Token: "mock_token"}

	// Inisialisasi mock JwtToken dan konfigurasinya
	suite.uum.On("GetByEmailPassword", "test@example.com", "password").Return(mockUser, nil)
	suite.jum.On("GenerateToken", mockUser).Return(mockToken, nil)

	// Membuat instance dari AuthUseCase menggunakan mock yang sudah disiapkan
	authUC := usecase.NewAuthUseCase(suite.uum, suite.jum)

	// Payload yang akan digunakan untuk Login
	payload := dto.AuthRequestDto{Email: "test@example.com", Password: "password"}

	// Melakukan Login dengan payload
	token, err := authUC.Login(payload)

	// Melakukan pengecekan hasil Login dengan ekspektasi
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockToken, token)

	// Memastikan bahwa metode-metode pada mock telah dipanggil sesuai ekspektasi
	suite.uum.AssertExpectations(suite.T())
	suite.jum.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLogin_GetByEmailPasswordError() {
	mockError := errors.New("some error")

	// Inisialisasi mock JwtToken dan konfigurasinya
	suite.uum.On("GetByEmailPassword", "test@example.com", "password").Return(entity.User{}, mockError)

	// Membuat instance dari AuthUseCase menggunakan mock yang sudah disiapkan
	authUC := usecase.NewAuthUseCase(suite.uum, suite.jum)

	// Payload yang akan digunakan untuk Login
	payload := dto.AuthRequestDto{Email: "test@example.com", Password: "password"}

	// Melakukan Login dengan payload
	_, err := authUC.Login(payload)

	// Melakukan pengecekan hasil Login dengan ekspektasi
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), mockError, err)

	// Memastikan bahwa metode-metode pada mock telah dipanggil sesuai ekspektasi
	suite.uum.AssertExpectations(suite.T())
	suite.jum.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLogin_GenerateTokenError() {
	mockUser := entity.User{Id: "1", Email: "test@example.com", Password: "password"}
	mockError := errors.New("token generation error")

	// Inisialisasi mock JwtToken dan konfigurasinya
	suite.uum.On("GetByEmailPassword", "test@example.com", "password").Return(mockUser, nil)
	suite.jum.On("GenerateToken", mockUser).Return(dto.AuthResponseDto{}, mockError)

	// Membuat instance dari AuthUseCase menggunakan mock yang sudah disiapkan
	authUC := usecase.NewAuthUseCase(suite.uum, suite.jum)

	// Payload yang akan digunakan untuk Login
	payload := dto.AuthRequestDto{Email: "test@example.com", Password: "password"}

	// Melakukan Login dengan payload
	_, err := authUC.Login(payload)

	// Melakukan pengecekan hasil Login dengan ekspektasi
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), mockError, err)

	// Memastikan bahwa metode-metode pada mock telah dipanggil sesuai ekspektasi
	suite.uum.AssertExpectations(suite.T())
	suite.jum.AssertExpectations(suite.T())
}
