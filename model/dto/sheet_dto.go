package dto

type SheetData struct {
	BookingId    string
	CheckIn      string
	CheckOut     string
	UserName     string
	CustomerName string
	IsAgree      bool
	Information  string
	TotalPrice   int
}

type GetBookingOneDayParams struct {
	Date string `json:"date"`
}

type GetBookingOneMonthParams struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type GetBookingOneYearParams struct {
	Year string `json:"year"`
}
