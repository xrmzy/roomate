package commonmock

import (
	"roomate/model/dto"

	"github.com/stretchr/testify/mock"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetCommonMock struct {
	mock.Mock
}

func (g *GoogleSheetCommonMock) NewService() (*sheets.Service, error) {
	args := g.Called()
	return args.Get(0).(*sheets.Service), args.Error(1)
}

func (g *GoogleSheetCommonMock) AppendSheet(sheetData []dto.SheetData, service *sheets.Service) error {
	args := g.Called(sheetData, service)
	return args.Error(0)
}

func (g *GoogleSheetCommonMock) DeleteSheetData(service *sheets.Service) error {
	args := g.Called(service)
	return args.Error(0)
}
