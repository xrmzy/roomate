package usecasemock

import (
	"roomate/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockGSheet struct {
	mock.Mock
}

func (m *MockGSheet) NewService() (*dto.SheetData, error) {
	args := m.Called()
	return args.Get(0).(*dto.SheetData), args.Error(1)
}

func (m *MockGSheet) AppendSheet(sheetData []dto.SheetData, service *dto.SheetData) error {
	args := m.Called(sheetData, service)
	return args.Error(0)
}

func (m *MockGSheet) DeleteSheetData(service *dto.SheetData) error {
	args := m.Called(service)
	return args.Error(0)
}
