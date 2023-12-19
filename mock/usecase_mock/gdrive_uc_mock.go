package usecasemock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockGDrive struct {
	mock.Mock
}

func (m *MockGDrive) NewService() (*http.Response, error) {
	args := m.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockGDrive) Download(service *http.Response) (*http.Response, error) {
	args := m.Called(service)
	return args.Get(0).(*http.Response), args.Error(1)
}
