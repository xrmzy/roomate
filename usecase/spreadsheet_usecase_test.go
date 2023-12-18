package usecase

// import (
// 	repomock "roomate/mock/repo_mock"
// 	usecasemock "roomate/mock/usecase_mock"
// 	"roomate/model/dto"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mocks untuk dependensi
// type MockBookingRepo struct {
// 	mock.Mock
// }

// func (m *MockBookingRepo) GetOneDay(date time.Time) (dto.SheetData, error) {
// 	args := m.Called(date)
// 	return args.Get(0).(dto.SheetData), args.Error(1)
// }

// func (m *MockBookingRepo) GetOneMonth(month string, year string) ([]dto.SheetData, error) {
// 	args := m.Called(month, year)
// 	return args.Get(0).([]dto.SheetData), args.Error(1)
// }

// func (m *MockBookingRepo) GetOneYear(year string) ([]dto.SheetData, error) {
// 	args := m.Called(year)
// 	return args.Get(0).([]dto.SheetData), args.Error(1)
// }

// // Mocking dependensi lainnya (UserUseCase, CustomerUseCase, GDrive, GSheet)

// func TestDailyReport_Success(t *testing.T) {
// 	// Buat mock untuk dependensi
// 	bookingRepoMock := new(repomock.BookingRepoMock)
// 	userUseCaseMock := new(usecasemock.MockUserUseCase)
// 	customerUseCaseMock := new(usecasemock.MockCustomerUseCase)
// 	gDriveMock := new(usecasemock.MockGDrive)
// 	gSheetMock := new(usecasemock.MockGSheet)
// 	// Buat mock dependensi lainnya (UserUseCase, CustomerUseCase, GDrive, GSheet)

// 	// Buat instance dari GSheetUseCase dengan dependensi yang di-mock
// 	gSheetUseCase := NewGSheetUseCase(bookingRepoMock, userUseCaseMock, customerUseCaseMock, gDriveMock, gSheetMock)

// 	// Persiapkan data uji dengan string
// 	dateString := time.Now().Format("2006-01-02") // Ubah time.Time menjadi string dengan format yang sesuai

// 	// Mock perilaku dari dependensi
// 	mockBooking := dto.SheetData{
// 		BookingId:    "123",
// 		CheckIn:      "2023-12-18T08:00:00Z",
// 		CheckOut:     "2023-12-20T10:00:00Z",
// 		UserName:     "John Doe",
// 		CustomerName: "Jane Smith",
// 		IsAgree:      true,
// 		Information:  "Sample information",
// 		TotalPrice:   250,
// 		Date:         dateString, // Menggunakan string yang sesuai dengan perubahan di DTO
// 	}
// 	bookingRepoMock.On("GetOneDay", time.Now()).Return(mockBooking, nil)
// 	// Mock perilaku dependensi lainnya

// 	// Panggil metode yang diuji
// 	resp, err := gSheetUseCase.DailyReport(dto.GetBookingOneDayParams{Date: dateString})

// 	// Penegasan
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	// Lakukan pengecekan lainnya sesuai kebutuhan

// 	// Periksa ekspektasi dari mock dependensi
// 	bookingRepoMock.AssertExpectations(t)
// 	// Periksa ekspektasi dependensi lainnya
// }

// func TestMonthlyReport_Success(t *testing.T) {
// 	// Buat mock untuk dependensi
// 	bookingRepoMock := new(MockBookingRepo)
// 	// Buat mock dependensi lainnya (UserUseCase, CustomerUseCase, GDrive, GSheet)

// 	// Buat instance dari GSheetUseCase dengan dependensi yang di-mock
// 	gSheetUseCase := NewGSheetUseCase(bookingRepoMock /* mock dependensi lainnya */)

// 	// Persiapkan data uji dengan string
// 	date := time.Now()
// 	month := date.Format("01")  // Ubah time.Time menjadi string dengan format yang sesuai
// 	year := date.Format("2006") // Ubah time.Time menjadi string dengan format yang sesuai

// 	// Mock perilaku dari dependensi
// 	mockBookings := []dto.SheetData{
// 		{
// 			BookingId:    "123",
// 			CheckIn:      "2023-12-18T08:00:00Z",
// 			CheckOut:     "2023-12-20T10:00:00Z",
// 			UserName:     "John Doe",
// 			CustomerName: "Jane Smith",
// 			IsAgree:      true,
// 			Information:  "Sample information",
// 			TotalPrice:   250,
// 			Date:         date.Format("2006-01-02"), // Ubah time.Time menjadi string dengan format yang sesuai
// 		},
// 		// Add more mock data if needed
// 	}
// 	bookingRepoMock.On("GetOneMonth", month, year).Return(mockBookings, nil)
// 	// Mock perilaku dependensi lainnya

// 	// Panggil metode yang diuji
// 	resp, err := gSheetUseCase.MonthlyReport(dto.GetBookingOneMonthParams{Month: month, Year: year})

// 	// Penegasan
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	// Lakukan pengecekan lainnya sesuai kebutuhan

// 	// Periksa ekspektasi dari mock dependensi
// 	bookingRepoMock.AssertExpectations(t)
// 	// Periksa ekspektasi dependensi lainnya
// }

// func TestYearlyReport_Success(t *testing.T) {
// 	// Buat mock untuk dependensi
// 	bookingRepoMock := new(MockBookingRepo)
// 	// Buat mock dependensi lainnya (UserUseCase, CustomerUseCase, GDrive, GSheet)

// 	// Buat instance dari GSheetUseCase dengan dependensi yang di-mock
// 	gSheetUseCase := NewGSheetUseCase(bookingRepoMock /* mock dependensi lainnya */)

// 	// Persiapkan data uji dengan string
// 	date := time.Now()
// 	year := date.Format("2006") // Ubah time.Time menjadi string dengan format yang sesuai

// 	// Mock perilaku dari dependensi
// 	mockBookings := []dto.SheetData{
// 		{
// 			BookingId:    "123",
// 			CheckIn:      "2023-12-18T08:00:00Z",
// 			CheckOut:     "2023-12-20T10:00:00Z",
// 			UserName:     "John Doe",
// 			CustomerName: "Jane Smith",
// 			IsAgree:      true,
// 			Information:  "Sample information",
// 			TotalPrice:   250,
// 			Date:         date.Format("2006-01-02"), // Ubah time.Time menjadi string dengan format yang sesuai
// 		},
// 		// Add more mock data if needed
// 	}
// 	bookingRepoMock.On("GetOneYear", year).Return(mockBookings, nil)
// 	// Mock perilaku dependensi lainnya

// 	// Panggil metode yang diuji
// 	resp, err := gSheetUseCase.YearlyReport(dto.GetBookingOneYearParams{Year: year})

// 	// Penegasan
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	// Lakukan pengecekan lainnya sesuai kebutuhan

// 	// Periksa ekspektasi dari mock dependensi
// 	bookingRepoMock.AssertExpectations(t)
// 	// Periksa ekspektasi dependensi lainnya
// }
