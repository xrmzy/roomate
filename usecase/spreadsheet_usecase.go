package usecase

import (
	"roomate/model/dto"
	"roomate/repository"
	"roomate/utils/common"
	"time"
)

type GSheetUseCase interface {
	DailyReport(payload dto.GetBookingOneDayParams) error
	MonthlyReport(payload dto.GetBookingOneMonthParams) error
	YearlyReport(payload dto.GetBookingOneYearParams) error
}

type gSheetUseCase struct {
	bookingRepo repository.BookingRepository
	userUc      UserUseCase
	customerUc  CustomerUseCase
	gDrive      common.GDrive
	gSheet      common.GSheet
}

func (s *gSheetUseCase) DailyReport(payload dto.GetBookingOneDayParams) error {
	// get booking data
	booking, err := s.bookingRepo.GetOneDay(payload.Date)
	if err != nil {
		return err
	}

	// get user name
	user, err := s.userUc.GetUser(booking.UserName) // booking.UserName masih user id
	if err != nil {
		return err
	}
	booking.UserName = user.Name

	// get customer name
	customer, err := s.customerUc.GetCustomer(booking.CustomerName) // booking.CustomerName masih id
	if err != nil {
		return err
	}
	booking.CustomerName = customer.Name

	// parse check in and check out
	parsedCheckIn, _ := time.Parse(time.RFC3339, booking.CheckIn)
	parsedCheckOut, _ := time.Parse(time.RFC3339, booking.CheckOut)
	booking.CheckIn = parsedCheckIn.Format("2006-01-02")
	booking.CheckOut = parsedCheckOut.Format("2006-01-02")

	// convert booking into slice
	bookingSlice := []dto.SheetData{booking}

	// get new sheet service
	service, err := s.gSheet.NewService()
	if err != nil {
		return err
	}

	// clear sheet data if exist
	err = s.gSheet.DeleteSheetData(service)
	if err != nil {
		return err
	}

	// append data to sheet
	err = s.gSheet.AppendSheet(bookingSlice, service)
	if err != nil {
		return err
	}

	// get new drive service
	driveService, err := s.gDrive.NewService()
	if err != nil {
		return err
	}

	// download sheet file
	err = s.gDrive.Download(driveService)
	if err != nil {
		return err
	}

	return nil
}

func (s *gSheetUseCase) MonthlyReport(payload dto.GetBookingOneMonthParams) error {
	bookings, err := s.bookingRepo.GetOneMonth(payload.Month, payload.Year)
	if err != nil {
		return err
	}

	var newBookings []dto.SheetData
	for _, booking := range bookings {
		// get user name
		user, err := s.userUc.GetUser(booking.UserName)
		if err != nil {
			return err
		}
		booking.UserName = user.Name

		// get customer name
		customer, err := s.customerUc.GetCustomer(booking.CustomerName)
		if err != nil {
			return err
		}
		booking.CustomerName = customer.Name

		// parse check in and check out
		parsedCheckIn, _ := time.Parse(time.RFC3339, booking.CheckIn)
		parsedCheckOut, _ := time.Parse(time.RFC3339, booking.CheckOut)
		booking.CheckIn = parsedCheckIn.Format("2006-01-02")
		booking.CheckOut = parsedCheckOut.Format("2006-01-02")

		newBookings = append(newBookings, booking)
	}

	// get new sheet service
	service, err := s.gSheet.NewService()
	if err != nil {
		return err
	}

	// clear sheet data if exist
	err = s.gSheet.DeleteSheetData(service)
	if err != nil {
		return err
	}

	// append data to sheet
	err = s.gSheet.AppendSheet(newBookings, service)
	if err != nil {
		return err
	}

	// get new drive service
	driveService, err := s.gDrive.NewService()
	if err != nil {
		return err
	}

	// download sheet file
	err = s.gDrive.Download(driveService)
	if err != nil {
		return err
	}

	return nil
}

func (s *gSheetUseCase) YearlyReport(payload dto.GetBookingOneYearParams) error {
	bookings, err := s.bookingRepo.GetOneYear(payload.Year)
	if err != nil {
		return err
	}

	var newBookings []dto.SheetData
	for _, booking := range bookings {
		// get user name
		user, err := s.userUc.GetUser(booking.UserName)
		if err != nil {
			return err
		}
		booking.UserName = user.Name

		// get customer name
		customer, err := s.customerUc.GetCustomer(booking.CustomerName)
		if err != nil {
			return err
		}
		booking.CustomerName = customer.Name

		// parse check in and check out
		parsedCheckIn, _ := time.Parse(time.RFC3339, booking.CheckIn)
		parsedCheckOut, _ := time.Parse(time.RFC3339, booking.CheckOut)
		booking.CheckIn = parsedCheckIn.Format("2006-01-02")
		booking.CheckOut = parsedCheckOut.Format("2006-01-02")

		newBookings = append(newBookings, booking)
	}

	// get new sheet service
	service, err := s.gSheet.NewService()
	if err != nil {
		return err
	}

	// clear sheet data if exist
	err = s.gSheet.DeleteSheetData(service)
	if err != nil {
		return err
	}

	// append data to sheet
	err = s.gSheet.AppendSheet(newBookings, service)
	if err != nil {
		return err
	}

	// get new drive service
	driveService, err := s.gDrive.NewService()
	if err != nil {
		return err
	}

	// download sheet file
	err = s.gDrive.Download(driveService)
	if err != nil {
		return err
	}

	return nil
}

func NewGSheetUseCase(
	bookingRepo repository.BookingRepository,
	userUc UserUseCase,
	customerUc CustomerUseCase,
	gDrive common.GDrive,
	gSheet common.GSheet,
) GSheetUseCase {
	return &gSheetUseCase{
		bookingRepo: bookingRepo,
		userUc:      userUc,
		customerUc:  customerUc,
		gDrive:      gDrive,
		gSheet:      gSheet,
	}
}
