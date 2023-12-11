package usecase

import (
	"roomate/model/dto"
	"roomate/model/entity"
	"roomate/repository"
	"time"
)

type BookingUsecase interface {
	GetAllBookings(payload dto.GetAllParams) ([]entity.Booking, error)
	GetBooking(id string) (entity.Booking, error)
	CreateBooking(payload dto.CreateBookingParams) (entity.Booking, error)
	// DeleteBooking(id string) error
}

type bookingUsecase struct {
	bookingRepo repository.BookingRepository
	userUc      UserUseCase
	customerUc  CustomerUseCase
	roomUc      RoomUseCase
	serviceUc   ServiceUseCase
}

func (u *bookingUsecase) GetAllBookings(payload dto.GetAllParams) ([]entity.Booking, error) {
	bookings, err := u.bookingRepo.GetAll(payload.Limit, payload.Offset)

	if err != nil {
		return bookings, err
	}

	return bookings, nil
}

func (u *bookingUsecase) GetBooking(id string) (entity.Booking, error) {
	booking, err := u.bookingRepo.Get(id)

	if err != nil {
		return booking, err
	}

	return booking, nil
}

// func (u *bookingUsecase) CreateBooking(payload dto.CreateBookingParams) (entity.Booking, error) {
// 	var booking entity.Booking
// 	var totalPrice int

// 	// parsing string into time
// 	layout := "2006-01-02"
// 	checkIn, _ := time.Parse(layout, payload.CheckIn)
// 	checkOut, _ := time.Parse(layout, payload.CheckOut)

// 	// calculate days
// 	duration := checkOut.Sub(checkIn)
// 	days := int(duration.Hours() / 24)
// 	booking.Night = days

// 	var bookingDetails []entity.BookingDetail
// 	// looping over booking details
// 	for _, detail := range payload.BookingDetails {
// 		var bookingDetail entity.BookingDetail

// 		// get room price
// 		room, err := u.roomUc.GetRoom(detail.RoomId)
// 		if err != nil {
// 			return booking, err
// 		}

// 		var totalServicesPrice int
// 		// looping booking detail services to get service price and name
// 		var services []entity.BookingDetailService
// 		for _, s := range detail.Services {
// 			var service entity.BookingDetailService

// 			// get service price
// 			serviceResult, err := u.serviceUc.GetService(s.ServiceId)
// 			if err != nil {
// 				return booking, err
// 			}

// 			service.ServiceId = s.ServiceId
// 			service.ServiceName = serviceResult.Name
// 			services = append(services, service)
// 			totalServicesPrice += serviceResult.Price
// 		}

// 		bookingDetail.RoomId = detail.RoomId
// 		bookingDetail.Services = services
// 		bookingDetail.SubTotal = totalServicesPrice + room.Price

// 		totalPrice += bookingDetail.SubTotal
// 		bookingDetails = append(bookingDetails, bookingDetail)
// 	}

// 	booking.CheckIn = checkIn
// 	booking.CheckOut = checkOut
// 	booking.UserId = payload.UserId
// 	booking.CustomerId = payload.CustomerId
// 	booking.BookingDetails = bookingDetails
// 	booking.TotalPrice = totalPrice

// 	booking, err := u.bookingRepo.Create(booking)

// 	if err != nil {
// 		return booking, err
// 	}

// 	return booking, nil
// }


func (u *bookingUsecase) CreateBooking(payload dto.CreateBookingParams) (entity.Booking, error) {
	// Initialize an empty Booking struct and totalPrice variable.
	booking := entity.Booking{}
	totalPrice := 0

	// Parse the check-in and check-out dates from the payload.
	checkIn, _ := time.Parse("2006-01-02", payload.CheckIn)
	checkOut, _ := time.Parse("2006-01-02", payload.CheckOut)

	// Calculate the number of nights by subtracting the check-in date from the check-out date.
	days := int(checkOut.Sub(checkIn).Hours() / 24)

	// Set the 'Night' field in the booking struct to the number of nights.
	booking.Night = days

	// Create an array of BookingDetail structs with the same length as the number of booking details in the payload.
	bookingDetails := make([]entity.BookingDetail, len(payload.BookingDetails))

	// Iterate over each booking detail in the payload.
	for i, detail := range payload.BookingDetails {
		// Get the room information for the current booking detail.
		room, err := u.roomUc.GetRoom(detail.RoomId)
		if err != nil {
			return booking, err
		}

		// Create an array of BookingDetailService structs with the same length as the number of services in the current booking detail.
		services := make([]entity.BookingDetailService, len(detail.Services))
		totalServicesPrice := 0

		// Iterate over each service in the current booking detail.
		for j, s := range detail.Services {
			// Get the service information for the current service.
			serviceResult, err := u.serviceUc.GetService(s.ServiceId)
			if err != nil {
				return booking, err
			}

			// Set the service ID and name in the BookingDetailService struct.
			services[j] = entity.BookingDetailService{
				ServiceId:   s.ServiceId,
				ServiceName: serviceResult.Name,
			}

			// Add the price of the current service to the total services price.
			totalServicesPrice += serviceResult.Price
		}

		// Set the room ID, services, and sub-total in the BookingDetail struct.
		bookingDetails[i] = entity.BookingDetail{
			RoomId:    detail.RoomId,
			Services:  services,
			SubTotal:  totalServicesPrice + room.Price,
		}

		// Add the sub-total of the current booking detail to the total price.
		totalPrice += bookingDetails[i].SubTotal
	}

	// Set the check-in, check-out, user ID, customer ID, booking details, and total price in the booking struct.
	booking.CheckIn = checkIn
	booking.CheckOut = checkOut
	booking.UserId = payload.UserId
	booking.CustomerId = payload.CustomerId
	booking.BookingDetails = bookingDetails
	booking.TotalPrice = totalPrice

	// Create the booking in the repository.
	booking, err := u.bookingRepo.Create(booking)
	if err != nil {
		return booking, err
	}

	// Return the created booking.
	return booking, nil
}

func NewBookingUseCase(
	bookingRepo repository.BookingRepository,
	userUc UserUseCase,
	customerUc CustomerUseCase,
	roomUc RoomUseCase,
	serviceUc ServiceUseCase,
) BookingUsecase {
	return &bookingUsecase{
		bookingRepo: bookingRepo,
		userUc:      userUc,
		customerUc:  customerUc,
		roomUc:      roomUc,
		serviceUc:   serviceUc,
	}
}
