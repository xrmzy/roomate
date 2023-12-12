package entity

import "time"

type Booking struct {
	Id             string          `json:"id"`
	Night          int             `json:"night"`
	CheckIn        time.Time       `json:"checkIn"`
	CheckOut       time.Time       `json:"checkOut"`
	UserId         string          `json:"userId"`
	CustomerId     string          `json:"customerId"`
	IsAgree        bool            `json:"isAgree"`
	Information    string          `json:"information"`
	BookingDetails []BookingDetail `json:"bookingDetails"`
	TotalPrice     int             `json:"totalPrice"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	IsDeleted      bool            `json:"isDeleted"`
}

type BookingDetail struct {
	Id        string                 `json:"id"`
	BookingId string                 `json:"bookingId"`
	RoomId    string                 `json:"roomId"`
	Services  []BookingDetailService `json:"services"`
	SubTotal  int                    `json:"subTotal"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	IsDeleted bool                   `json:"isDeleted"`
}

type BookingDetailService struct {
	Id              string    `json:"id"`
	BookingDetailId string    `json:"bookingDetailId"`
	ServiceId       string    `json:"serviceId"`
	ServiceName     string    `json:"serviceName"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	IsDeleted       bool      `json:"isDeleted"`
}
