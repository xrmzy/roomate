package usecase

import (
	"net/http"
	commonmock "roomate/mock/common_mock"
	repomock "roomate/mock/repo_mock"
	usecasemock "roomate/mock/usecase_mock"
	"roomate/model/dto"
	"roomate/model/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
)

type GSheetUseCaseTestSuite struct {
	suite.Suite
	brm *repomock.BookingRepoMock
	uum *usecasemock.UserUseCaseMock
	cum *usecasemock.CustomerUseCaseMock
	gdm *commonmock.GoogleDriveCommonMock
	gsm *commonmock.GoogleSheetCommonMock
	gs  GSheetUseCase
}

func (suite *GSheetUseCaseTestSuite) SetupTest() {
	suite.brm = new(repomock.BookingRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.cum = new(usecasemock.CustomerUseCaseMock)
	suite.gdm = new(commonmock.GoogleDriveCommonMock)
	suite.gsm = new(commonmock.GoogleSheetCommonMock)
	suite.gs = NewGSheetUseCase(suite.brm, suite.uum, suite.cum, suite.gdm, suite.gsm)
}

var (
	oneDayParam = dto.GetBookingOneDayParams{
		Date: "2023-12-14",
	}
	// day   = "2023-12-14"
	// month = "12"
	// year  = "2023"

	dummySheetData = dto.SheetData{
		BookingId:    "1",
		CheckIn:      "2023-12-14",
		CheckOut:     "2023-12-16",
		UserName:     "John",
		CustomerName: "Jane",
		IsAgree:      true,
		Information:  "Accepted",
		TotalPrice:   1000,
	}

	dummyUser = entity.User{
		Id:        "1",
		Name:      "John",
		Email:     "johndoe@me.com",
		Password:  "167916",
		RoleId:    "1",
		RoleName:  "admin",
		UpdatedAt: time.Now().Truncate(time.Second),
	}
)

func TestGSheetUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GSheetUseCaseTestSuite))
}

func (suite *GSheetUseCaseTestSuite) TestDailyReport() {
	var service *sheets.Service
	var driveService *drive.Service
	var resp *http.Response
	// get booking data
	suite.brm.On("GetOneDay", oneDayParam.Date).Return(dummySheetData, nil).Once()
	booking, err := suite.brm.GetOneDay(oneDayParam.Date)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), booking.BookingId, dummySheetData.BookingId)

	// get user data
	suite.uum.On("GetUser", dummyBooking.UserId).Return(dummyUser, nil).Once()
	user, err := suite.uum.GetUser(dummyBooking.UserId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Name, dummyUser.Name)

	// get customer data
	suite.cum.On("GetCustomer", dummyBooking.CustomerId).Return(dummyCustomer, nil).Once()
	customer, err := suite.cum.GetCustomer(dummyBooking.CustomerId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), customer.Name, dummyCustomer.Name)

	// convert booking into slice
	sheetDataSlice := []dto.SheetData{dummySheetData}

	// get sheet service
	suite.gsm.On("NewService").Return(service, nil)
	_, err = suite.gsm.NewService()
	assert.NoError(suite.T(), err)

	// clear sheet data
	suite.gsm.On("DeleteSheetData", service).Return(nil).Once()
	err = suite.gsm.DeleteSheetData(service)
	assert.NoError(suite.T(), err)

	// write sheet data
	suite.gsm.On("AppendSheet", sheetDataSlice, service).Return(nil).Once()
	err = suite.gsm.AppendSheet(sheetDataSlice, service)
	assert.NoError(suite.T(), err)

	// get drive service
	suite.gdm.On("NewService").Return(driveService, nil).Once()
	_, err = suite.gdm.NewService()
	assert.NoError(suite.T(), err)

	// download sheet file
	suite.gdm.On("Download", driveService).Return(resp, nil).Once()
	_, err = suite.gdm.Download(driveService)
	assert.NoError(suite.T(), err)
}
