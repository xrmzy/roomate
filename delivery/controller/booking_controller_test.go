package controller

import (
	"net/http"
	"net/http/httptest"
	middlewaremock "roomate/mock/middleware_mock"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingControllerTestSuite struct {
	suite.Suite
	bum                *usecasemock.BookingUseCaseMock
	engine             *gin.Engine
	amm                *middlewaremock.AuthMiddlewareMock
	expectedDateFormat string
}

func (suite *BookingControllerTestSuite) SetupTest() {
	suite.bum = new(usecasemock.BookingUseCaseMock)
	suite.engine = gin.Default()
	suite.amm = new(middlewaremock.AuthMiddlewareMock)
	suite.expectedDateFormat = "02-01-2006"

}

func TestBookingControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BookingControllerTestSuite))
}

var dummyBooking = entity.Booking{
	Id:          "1",
	Night:       3,
	CheckIn:     time.Now(),
	CheckOut:    time.Now().Add(3 * 24 * time.Hour),
	UserId:      "1",
	CustomerId:  "1",
	IsAgree:     true,
	Information: "Available",
	BookingDetails: []entity.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			RoomId:    "1",
			Services: []entity.BookingDetailService{
				{
					Id:              "1",
					BookingDetailId: "1",
					ServiceId:       "1",
					ServiceName:     "Proyektor",
				},
			},
			SubTotal:  200000,
			IsDeleted: false,
		},
	},
	TotalPrice: 300000,
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
	IsDeleted:  false,
}

var mockPayload = dto.CreateBookingParams{
	CheckIn:    "20-12-2023",
	CheckOut:   "23-12-2023",
	UserId:     "1",
	CustomerId: "1",
	BookingDetails: []entity.BookingDetail{
		{
			Services: []entity.BookingDetailService{
				{
					Id: "1",
				},
			},
		},
	},
}

var mockTokenJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI3OTExMzcsImlhdCI6MTcwMjcwNDczNywidXNlcklkIjoiZmI5ZWJkYTUtZjE5MS00MzA0LTljZWYtODdiOWQyOWM3YjY4Iiwicm9sZSI6IkFkbWluIn0.jg71NzO_UCb2MQXE_R12AnM_gwmEP9n7VkktL1Y2hxs"

func (suite *BookingControllerTestSuite) TestCreateBookingHandler_Failure() {
	bookingController := NewBookingController(suite.bum, suite.engine.Group("/api/v1"), suite.amm)
	bookingController.Route()

	//EKSPETASI
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bookings", nil)
	assert.NoError(suite.T(), err)

	// EKSEKUSI
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", dummyBooking.UserId)
	bookingController.CreateHandler(ctx)

	// ASSERTION
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

// fun
func (suite *BookingControllerTestSuite) TestCreateBookingHandlerCustomer_Failure() {
	bookingController := NewBookingController(suite.bum, suite.engine.Group("/api/v1"), suite.amm)
	bookingController.Route()

	//EKSPETASI
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bookings", nil)
	assert.NoError(suite.T(), err)

	// EKSEKUSI
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("customer", dummyBooking.CustomerId)
	bookingController.CreateHandler(ctx)

	// ASSERTION
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}
