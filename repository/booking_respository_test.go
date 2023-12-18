package repository

import (
	"database/sql"
	"errors"

	"regexp"
	"roomate/model/entity"
	"roomate/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// create suite
type BookingRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    BookingRepository
}

// setup
func (suite *BookingRepoTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.sqlmock = sqlmock
	suite.repo = NewBookingRepository(suite.mockDB)
}

func TestBookingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookingRepoTestSuite))
}

// dummy booking
var (
	checkIn, _   = time.Parse("2006-01-02", "2023-12-14")
	checkOut, _  = time.Parse("2006-01-02", "2023-12-16")
	dummyBooking = entity.Booking{
		Id:          "1",
		Night:       2,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
		UserId:      "1",
		CustomerId:  "1",
		IsAgree:     false,
		Information: "Pending",
		BookingDetails: []entity.BookingDetail{
			{
				Id:        "22",
				BookingId: "1",
				RoomId:    "1",
				Services: []entity.BookingDetailService{
					{
						Id:              "55",
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

func (suite *BookingRepoTestSuite) TestRepository_CreateBooking() {
	// test fail
	suite.sqlmock.ExpectBegin().WillReturnError(errors.New("Error begin"))
	_, err := suite.repo.Create(dummyBooking)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "Error begin", err.Error())

	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateBooking)).WithArgs(dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.TotalPrice, dummyBooking.UpdatedAt).WillReturnError(errors.New("insert failed"))

	_, err = suite.repo.Create(dummyBooking)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "insert failed", err.Error())

	// expectation
	suite.sqlmock.ExpectBegin()

	// design returning rows booking
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateBooking)).WithArgs(dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.TotalPrice, dummyBooking.UpdatedAt).WillReturnRows(rows)

	for _, i := range dummyBooking.BookingDetails {
		// design returning rows booking details
		rows := sqlmock.NewRows([]string{
			"id", "booking_id", "room_id", "sub_total", "created_at", "updated_at",
		}).AddRow(i.Id, i.BookingId, i.RoomId, i.SubTotal, i.CreatedAt, i.UpdatedAt)
		suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateBookingDetail)).WithArgs(i.BookingId, i.RoomId, i.SubTotal, i.UpdatedAt).WillReturnRows(rows)

		for _, j := range i.Services {
			// design returning rows booking detail service
			rows := sqlmock.NewRows([]string{
				"id", "booking_detail_id", "service_id", "service_name", "created_at", "updated_at",
			}).AddRow(j.Id, j.BookingDetailId, j.ServiceId, j.ServiceName, j.CreatedAt, j.UpdatedAt)
			suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.CreateBookingDetailService)).WithArgs(i.Id, j.ServiceId, j.ServiceName, j.UpdatedAt).WillReturnRows(rows)
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

// test get booking
func (suite *BookingRepoTestSuite) TestRepository_GetBooking() {
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetBooking)).WithArgs(dummyBooking.Id).WillReturnError(errors.New("get booking failed"))

	_, err := suite.repo.Get(dummyBooking.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "get booking failed", err.Error())

	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetBooking)).WithArgs(dummyBooking.Id).WillReturnRows(rows)

	for _, i := range dummyBooking.BookingDetails {
		detailRows := sqlmock.NewRows([]string{
			"id", "booking_id", "room_id", "sub_total", "created_at", "updated_at",
		}).AddRow(i.Id, i.BookingId, i.RoomId, i.SubTotal, i.CreatedAt, i.UpdatedAt)

		suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookingDetails)).WithArgs(i.BookingId).WillReturnRows(detailRows)

		for _, j := range i.Services {
			serviceRows := sqlmock.NewRows([]string{
				"id", "booking_detail_id", "service_id", "service_name", "created_at", "updated_at",
			}).AddRow(j.Id, j.BookingDetailId, j.ServiceId, j.ServiceName, j.CreatedAt, j.UpdatedAt)

			suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookingDetailServices)).WithArgs(i.Id).WillReturnRows(serviceRows)
		}
	}

	actual, err := suite.repo.Get(dummyBooking.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), dummyBooking.Id, actual.Id)
	assert.Equal(suite.T(), dummyBooking.BookingDetails[0].BookingId, actual.BookingDetails[0].BookingId)

}

func (suite *BookingRepoTestSuite) TestRepository_GetAllBookings() {
	// test failed
	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookings)).WithArgs(1, 0).WillReturnError(sql.ErrNoRows)
	_, err := suite.repo.GetAll(1, 0)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, dummyBooking.IsAgree, dummyBooking.Information, dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookings)).WithArgs(1, 0).WillReturnRows(rows)

	for _, i := range dummyBooking.BookingDetails {
		detailRows := sqlmock.NewRows([]string{
			"id", "booking_id", "room_id", "sub_total", "created_at", "updated_at",
		}).AddRow(i.Id, i.BookingId, i.RoomId, i.SubTotal, i.CreatedAt, i.UpdatedAt)

		suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookingDetails)).WithArgs(i.BookingId).WillReturnRows(detailRows)

		for _, j := range i.Services {
			serviceRows := sqlmock.NewRows([]string{
				"id", "booking_detail_id", "service_id", "service_name", "created_at", "updated_at",
			}).AddRow(j.Id, j.BookingDetailId, j.ServiceId, j.ServiceName, j.CreatedAt, j.UpdatedAt)

			suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.GetAllBookingDetailServices)).WithArgs(i.Id).WillReturnRows(serviceRows)
		}
	}

	actual, err := suite.repo.GetAll(1, 0)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyBooking.Id, actual[0].Id)
	assert.Equal(suite.T(), dummyBooking.BookingDetails[0].BookingId, actual[0].BookingDetails[0].BookingId)
}

// test update status booking
func (suite *BookingRepoTestSuite) TestUpdateStatusBooking() {
	// test failed
	suite.sqlmock.ExpectQuery("UPDATE bookings").WithArgs(dummyBooking.Id, true, "Accepted").WillReturnError(sql.ErrNoRows)

	_, err := suite.repo.UpdateStatus(dummyBooking.Id, true, "Accepted")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), sql.ErrNoRows, err)

	// test success
	rows := sqlmock.NewRows([]string{
		"id", "night", "check_in", "check_out", "user_id", "customer_id", "is_agree", "information", "total_price", "created_at", "updated_at",
	}).AddRow(dummyBooking.Id, dummyBooking.Night, dummyBooking.CheckIn, dummyBooking.CheckOut, dummyBooking.UserId, dummyBooking.CustomerId, true, "Accepted", dummyBooking.TotalPrice, dummyBooking.CreatedAt, dummyBooking.UpdatedAt)

	suite.sqlmock.ExpectQuery(regexp.QuoteMeta(common.UpdateBookingStatus)).WithArgs(dummyBooking.Id, true, "Accepted").WillReturnRows(rows)

	actual, err := suite.repo.UpdateStatus(dummyBooking.Id, true, "Accepted")
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), dummyBooking.IsAgree, actual.IsAgree)
	assert.NotEqual(suite.T(), dummyBooking.Information, actual.Information)

}

func (suite *BookingRepoTestSuite) TestRepository_DeleteBooking() {
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteBooking)).WithArgs(dummyBooking.Id).
		WillReturnError(errors.New("delete failed"))

	err := suite.repo.Delete(dummyBooking.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete failed", err.Error())

	// test success
	suite.sqlmock.ExpectExec(regexp.QuoteMeta(common.DeleteBooking)).WithArgs(dummyBooking.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = suite.repo.Delete(dummyBooking.Id)

	assert.NoError(suite.T(), err)
}
