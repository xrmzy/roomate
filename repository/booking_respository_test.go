package repository

import (
	"database/sql"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// create suite
type BookingRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    BookingRepository
}

// setup
func (suite *BookingRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewBookingRepository(suite.mockDB)
}

func TestBookingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookingRepositoryTestSuite))
}

// dummy booking
var (
	checkIn, _   = time.Parse("2006-01-02", "2023-12-14")
	CheckOut, _  = time.Parse("2006-01-02", "2023-12-16")
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
				SubTotal:  150000,
			},
		},
		TotalPrice: 150000,
		UpdatedAt:  time.Now().Truncate(time.Second),
		IsDeleted:  false,
	}
)

// test create booking success
func (suite *BookingRepositoryTestSuite) TestCreateBooking_Success() {
	// expectation
	suite.sqlmock.ExpectBegin()

	// design returning rows booking
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at", "is_deleted",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt, dummyBooking.IsDeleted)

	suite.sqlmock.ExpectQuery("INSERT INTO bookings").WithArgs(dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.TotalPrice, dummyBooking.UpdatedAt).WillReturnRows(rows)

	for _, i := range dummyBooking.BookingDetails {
		// design returning rows booking details
		rows := sqlmock.NewRows([]string{
			"id", "booking_id", "room_id", "sub_total", "created_at", "updated_at", "is_deleted",
		}).AddRow(i.Id, i.BookingId, i.RoomId, i.SubTotal, i.CreatedAt, i.UpdatedAt, i.IsDeleted)
		suite.sqlmock.ExpectQuery("INSERT INTO booking_details").WithArgs(i.BookingId, i.RoomId, i.SubTotal, i.UpdatedAt).WillReturnRows(rows)

		for _, j := range i.Services {
			// design returning rows booking detail service
			rows := sqlmock.NewRows([]string{
				"id", "booking_detail_id", "service_id", "service_name", "created_at", "updated_at", "is_deleted",
			}).AddRow(j.Id, j.BookingDetailId, j.ServiceId, j.ServiceName, j.CreatedAt, j.UpdatedAt, j.IsDeleted)
			suite.sqlmock.ExpectQuery("INSERT INTO booking_detail_services").WithArgs(i.Id, j.ServiceId, j.ServiceName, j.UpdatedAt).WillReturnRows(rows)
		}
	}

	suite.sqlmock.ExpectCommit()

	// actual
	actual, err := suite.repo.Create(dummyBooking)

	// assertion
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyBooking.UserId, actual.UserId)
}

// test get booking success
func (suite *BookingRepositoryTestSuite) TestGetBooking_Success() {
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at", "is_deleted",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt, dummyBooking.IsDeleted)

	suite.sqlmock.ExpectQuery("SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings").WithArgs(dummyBooking.Id).WillReturnRows(rows)

	for _, i := range dummyBooking.BookingDetails {
		rows := sqlmock.NewRows([]string{
			"id", "booking_id", "room_id", "sub_total", "created_at", "updated_at", "is_deleted",
		}).AddRow(i.Id, i.BookingId, i.RoomId, i.SubTotal, i.CreatedAt, i.UpdatedAt, i.IsDeleted)
		suite.sqlmock.ExpectQuery("SELECT id, booking_id, room_id, sub_total, created_at, updated_at, is_deleted FROM booking_details").WithArgs(dummyBooking.Id).WillReturnRows(rows)

		for _, j := range i.Services {
			rows := sqlmock.NewRows([]string{
				"id", "booking_detail_id", "service_id", "service_name", "created_at", "updated_at", "is_deleted",
			}).AddRow(j.Id, j.BookingDetailId, j.ServiceId, j.ServiceName, j.CreatedAt, j.UpdatedAt, j.IsDeleted)
			suite.sqlmock.ExpectQuery("SELECT id, booking_detail_id, service_id, service_name, created_at, updated_at, is_deleted FROM booking_detail_services").WithArgs(i.Id).WillReturnRows(rows)
		}
	}

	actual, err := suite.repo.Get(dummyBooking.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyBooking.Id, actual.Id)
	assert.Equal(suite.T(), dummyBooking.BookingDetails[0].Id, actual.BookingDetails[0].Id)
	assert.Equal(suite.T(), dummyBooking.BookingDetails[0].Services[0].Id, actual.BookingDetails[0].Services[0].Id)
}

// test get all bookings
func (suite *BookingRepositoryTestSuite) TestGetAllBookings() {
	// test failed
	rowsErr := sqlmock.NewRows([]string{
		"night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at", "is_deleted",
	}).AddRow(dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt, dummyBooking.IsDeleted)

	suite.sqlmock.ExpectQuery("SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings").WithArgs(1, 0).WillReturnError(sql.ErrNoRows)
	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	suite.sqlmock.ExpectQuery("SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings").WithArgs(1, 0).WillReturnRows(rowsErr)
	_, err = suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)

	// test success
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at", "is_deleted",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt, dummyBooking.IsDeleted)

	suite.sqlmock.ExpectQuery("SELECT id, night, check_in, check_out, user_id, customer_id, is_agree, information, total_price, created_at, updated_at, is_deleted FROM bookings").WithArgs(1, 0).WillReturnRows(rows)

	actual, err := suite.repo.GetAll(1, 0)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyBooking.Id, actual[0].Id)
}

// test update status booking
func (suite *BookingRepositoryTestSuite) TestUpdateStatusBooking() {
	// test failed
	suite.sqlmock.ExpectQuery("UPDATE bookings").WithArgs(dummyBooking.Id, true, "Accepted").WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.UpdateStatus(dummyBooking.Id, true, "Accepted")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at", "is_deleted",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, true, "Accepted", dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt, dummyBooking.IsDeleted)

	suite.sqlmock.ExpectQuery("UPDATE bookings").WithArgs(dummyBooking.Id, true, "Accepted").WillReturnRows(rows)

	actual, err := suite.repo.UpdateStatus(dummyBooking.Id, true, "Accepted")
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), dummyBooking.IsAgree, actual.IsAgree)
	assert.NotEqual(suite.T(), dummyBooking.Information, actual.Information)

}
