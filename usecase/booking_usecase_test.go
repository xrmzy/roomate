package usecase

import (
	"errors"
	repomock "roomate/mock/repo_mock"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingUseCaseTestSuite struct {
	suite.Suite
	brm *repomock.BookingRepoMock
	rum *usecasemock.RoomUseCaseMock
	sum *usecasemock.ServiceUseCaseMock
	bu  BookingUsecase
}

func (suite *BookingUseCaseTestSuite) SetupTest() {
	suite.brm = new(repomock.BookingRepoMock)
	suite.rum = new(usecasemock.RoomUseCaseMock)
	suite.sum = new(usecasemock.ServiceUseCaseMock)
	suite.bu = NewBookingUseCase(suite.brm, suite.rum, suite.sum)
}

func TestBookingUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookingUseCaseTestSuite))
}

var (
	checkIn, _  = time.Parse("2006-01-02", "2023-12-14")
	CheckOut, _ = time.Parse("2006-01-02", "2023-12-16")

	dummyBooking = entity.Booking{
		Id:          "1",
		Night:       2,
		CheckIn:     checkIn,
		CheckOut:    CheckOut,
		UserId:      "1",
		CustomerId:  "1",
		IsAgree:     false,
		Information: "Pending",
		BookingDetails: []entity.BookingDetail{
			{
				BookingId: "1",
				RoomId:    "1",
				Services: []entity.BookingDetailService{
					{
						BookingDetailId: "1",
						ServiceId:       "1",
						ServiceName:     "Dolby Atmos 8D Surround Audio",
						UpdatedAt:       time.Now().Truncate(time.Second),
					},
				},
				UpdatedAt: time.Now().Truncate(time.Second),
				SubTotal:  380000,
			},
		},
		TotalPrice: 380000,
		UpdatedAt:  time.Now().Truncate(time.Second),
		IsDeleted:  false,
	}

	dummyRoom = entity.Room{
		Id:         "1",
		RoomNumber: "101",
		RoomType:   "Deluxe",
		Capacity:   2,
		Facility:   "Private Pool",
		Price:      150000,
		Status:     "Available",
	}

	dummyService = entity.Service{
		Id:    "1",
		Name:  "Dolby Atmos 8D Surround Audio",
		Price: 80000,
	}

	dummyPayload = dto.CreateBookingParams{
		CheckIn:    "2023-12-14",
		CheckOut:   "2023-12-16",
		UserId:     "1",
		CustomerId: "1",
		BookingDetails: []entity.BookingDetail{
			{
				RoomId: "1",
				Services: []entity.BookingDetailService{
					{
						ServiceId: "1",
					},
				},
			},
		},
	}

	dummyGetAllParams = dto.GetAllParams{
		Limit:  10,
		Offset: 0,
	}

	dummyUpdateParams = dto.UpdateBookingStatusParams{
		BookingId:   "1",
		IsAgree:     false,
		Information: "Pending",
	}
)

func (suite *BookingUseCaseTestSuite) TestCreateBooking() {
	booking := entity.Booking{}
	var totalPrice int

	checkIn, _ := time.Parse("2006-01-02", dummyPayload.CheckIn)
	checkOut, _ := time.Parse("2006-01-02", dummyPayload.CheckOut)

	booking.Night = int(checkOut.Sub(checkIn).Hours() / 24)

	var bookingDetails []entity.BookingDetail
	for _, detail := range dummyPayload.BookingDetails {
		suite.rum.On("GetRoom", detail.RoomId).Return(dummyRoom, nil).Once()

		var services []entity.BookingDetailService
		var totalServicePrice int
		for _, service := range detail.Services {
			suite.sum.On("GetService", service.ServiceId).Return(dummyService, nil).Once()

			services = append(services, entity.BookingDetailService{
				ServiceId:   dummyService.Id,
				ServiceName: dummyService.Name,
			})

			totalServicePrice += dummyService.Price
		}

		bookingDetails = append(bookingDetails, entity.BookingDetail{
			RoomId:   detail.RoomId,
			Services: services,
			SubTotal: totalServicePrice + dummyRoom.Price*booking.Night,
		})

		totalPrice += bookingDetails[0].SubTotal
	}

	booking.CheckIn = checkIn
	booking.CheckOut = checkOut
	booking.UserId = dummyPayload.UserId
	booking.CustomerId = dummyPayload.CustomerId
	booking.BookingDetails = bookingDetails
	booking.TotalPrice = totalPrice

	// create booking success
	suite.brm.On("Create", booking).Return(booking, nil).Once()
	_, err := suite.bu.CreateBooking(dummyPayload)

	suite.Require().NoError(err)
	assert.Equal(suite.T(), booking.TotalPrice, dummyBooking.TotalPrice)
}

func (suite *BookingUseCaseTestSuite) TestGetBooking() {
	// get booking fail
	suite.brm.On("Get", dummyBooking.Id).Return(entity.Booking{}, errors.New("booking not found")).Once()

	bookingErr, err := suite.bu.GetBooking(dummyBooking.Id)
	suite.Require().Error(err)
	suite.Require().EqualError(err, "booking not found")
	suite.Require().Equal(entity.Booking{}, bookingErr)

	// get booking success
	suite.brm.On("Get", dummyBooking.Id).Return(dummyBooking, nil).Once()

	booking, err := suite.bu.GetBooking(dummyBooking.Id)

	suite.Require().NoError(err)
	suite.Require().Equal(dummyBooking, booking)
}

func (suite *BookingUseCaseTestSuite) TestGetAllBooking() {
	// get all booking fail
	var limit int = 10
	var offset int
	suite.brm.On("GetAll", limit, offset).Return([]entity.Booking{}, errors.New("booking not found")).Once()

	bookingErr, err := suite.bu.GetAllBookings(dummyGetAllParams)
	suite.Require().Error(err)
	suite.Require().EqualError(err, "booking not found")
	suite.Require().Equal([]entity.Booking{}, bookingErr)

	// get all booking success
	suite.brm.On("GetAll", limit, offset).Return([]entity.Booking{dummyBooking}, nil).Once()

	booking, err := suite.bu.GetAllBookings(dummyGetAllParams)

	suite.Require().NoError(err)
	suite.Require().Equal([]entity.Booking{dummyBooking}, booking)
}

func (suite *BookingUseCaseTestSuite) TestUpdateBookingStatus() {
	// update booking status fail
	suite.brm.On("UpdateStatus", dummyBooking.Id, dummyBooking.IsAgree, dummyBooking.Information).Return(entity.Booking{}, errors.New("booking not found")).Once()

	bookingErr, err := suite.bu.UpdateBookingStatus(dummyUpdateParams)
	suite.Require().Error(err)
	suite.Require().EqualError(err, "booking not found")
	suite.Require().Equal(entity.Booking{}, bookingErr)

	// update booking status success
	suite.brm.On("UpdateStatus", dummyBooking.Id, dummyBooking.IsAgree, dummyBooking.Information).Return(dummyBooking, nil).Once()

	booking, err := suite.bu.UpdateBookingStatus(dummyUpdateParams)
	suite.Require().NoError(err)
	suite.Require().Equal(dummyBooking, booking)
}
