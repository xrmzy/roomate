package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type MockJwtToken struct {
	mock.Mock
}

// GenerateToken implements common.JwtToken.
func (m *MockJwtToken) GenerateToken(payload entity.User) (dto.AuthResponseDto, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

// VerifyToken implements common.JwtToken.
func (m *MockJwtToken) VerifyToken(token string) (jwt.MapClaims, error) {
	args := m.Called(token)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

// RefreshToken implements common.JwtToken.
func (m *MockJwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	args := m.Called(oldTokenString)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func NewMockJwtToken() *MockJwtToken {
	return &MockJwtToken{}
}
