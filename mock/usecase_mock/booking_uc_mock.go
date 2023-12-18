package usecasemock

import (
	"roomate/model/dto"
	"roomate/model/entity"

	"github.com/stretchr/testify/mock"
)

// BookingUseCaseMock adalah mock untuk BookingUsecase
type BookingUseCaseMock struct {
	mock.Mock
}

// GetAllBookings adalah metode mock untuk GetAllBookings di BookingUsecase
func (m *BookingUseCaseMock) GetAllBookings(payload dto.GetAllParams) ([]entity.Booking, error) {
	args := m.Called(payload)
	return args.Get(0).([]entity.Booking), args.Error(1)
}

// GetBooking adalah metode mock untuk GetBooking di BookingUsecase
func (m *BookingUseCaseMock) GetBooking(id string) (entity.Booking, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Booking), args.Error(1)
}

// CreateBooking adalah metode mock untuk CreateBooking di BookingUsecase
func (m *BookingUseCaseMock) CreateBooking(payload dto.CreateBookingParams) (entity.Booking, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Booking), args.Error(1)
}

// UpdateBookingStatus adalah metode mock untuk UpdateBookingStatus di BookingUsecase
func (m *BookingUseCaseMock) UpdateBookingStatus(payload dto.UpdateBookingStatusParams) (entity.Booking, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Booking), args.Error(1)
}
