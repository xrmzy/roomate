package repomock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

type BookingRepoMock struct {
	mock.Mock
}

func (b *BookingRepoMock) Create(booking entity.Booking) (entity.Booking, error) {
	args := b.Called(booking)
	return args.Get(0).(entity.Booking), args.Error(1)
}

func (b *BookingRepoMock) GetAll(limit, offset int) ([]entity.Booking, error) {
	args := b.Called(limit, offset)
	return args.Get(0).([]entity.Booking), args.Error(1)
}

func (b *BookingRepoMock) Get(id string) (entity.Booking, error) {
	args := b.Called(id)
	return args.Get(0).(entity.Booking), args.Error(1)
}

func (b *BookingRepoMock) UpdateStatus(id string, isAgree bool, information string) (entity.Booking, error) {
	args := b.Called(id, isAgree, information)
	return args.Get(0).(entity.Booking), args.Error(1)
}

func (b *BookingRepoMock) Delete(id string) error {
	args := b.Called(id)
	return args.Error(0)
}

func (b *BookingRepoMock) GetOneDay(date string) (dto.SheetData, error) {
	args := b.Called(date)
	return args.Get(0).(dto.SheetData), args.Error(1)
}

func (b *BookingRepoMock) GetOneMonth(month, year string) ([]dto.SheetData, error) {
	args := b.Called(month, year)
	return args.Get(0).([]dto.SheetData), args.Error(1)
}

func (b *BookingRepoMock) GetOneYear(year string) ([]dto.SheetData, error) {
	args := b.Called(year)
	return args.Get(0).([]dto.SheetData), args.Error(1)
}
