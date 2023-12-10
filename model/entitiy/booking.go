package entity

import "time"

type Booking struct {
	ID             string          `json:"id"`
	Night          int             `json:"night"`
	CheckIn        string          `json:"checkIn"`
	CheckOut       string          `json:"checkOut"`
	UserID         string          `json:"userId"`
	CustomerID     string          `json:"customerId"`
	CustomerName   string          `json:"customerName"`
	Status         bool            `json:"status"`
	BookingDetails []BookingDetail `json:"bookingDetail"`
	Information    string          `json:"information"`
	TotalPrice     int64           `json:"totalPrice"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpadatedAt     time.Time       `json:"updatedAt"`
	IsDeleted      bool            `json:"isDeleted"`
}

type BookingDetail struct {
	Id        string    `json:"id"`
	BookingID string    `json:"bookingId"`
	RoomID    string    `json:"roomId"`
	ServiceID Service   `json:"serviceId"`
	SubTotal  int64     `json:"subTotal"`
	CreatedAt time.Time `json:"createdAt"`
	UpadtedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
