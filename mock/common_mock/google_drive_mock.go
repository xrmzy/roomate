package commonmock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
	"google.golang.org/api/drive/v3"
)

type GoogleDriveCommonMock struct {
	mock.Mock
}

func (g *GoogleDriveCommonMock) NewService() (*drive.Service, error) {
	args := g.Called()
	return args.Get(0).(*drive.Service), args.Error(1)
}

func (g *GoogleDriveCommonMock) Download(service *drive.Service) (*http.Response, error) {
	args := g.Called(service)
	return args.Get(0).(*http.Response), args.Error(1)
}
