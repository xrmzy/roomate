package usecasemock

import (
	"roomate/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockAuthUseCase struct {
	mock.Mock
}

// Login adalah implementasi mock untuk Login di AuthUseCase.
func (m *MockAuthUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

// NewMockAuthUseCase membuat instance baru dari MockAuthUseCase.
func NewMockAuthUseCase() *MockAuthUseCase {
	return &MockAuthUseCase{}
}
