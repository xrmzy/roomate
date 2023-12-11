package dto

import (
	"roomate/model/entity"
)

type CreateBookingParams struct {
	CheckIn        string                 `json:"checkIn"`
	CheckOut       string                 `json:"checkOut"`
	UserId         string                 `json:"userId"`
	CustomerId     string                 `json:"customerId"`
	BookingDetails []entity.BookingDetail `json:"bookingDetails"`
}

// contoh request json
/*
	{
 	"checkIn": "2022-01-01",
 	"checkOut": "2022-01-02",
	"userId": "1",
	"customerId": "1",
	"bookingDetails": [
		{
			"roomId": "1",
			"services": [
				{
					"serviceId": "1",
				},
				{
					"serviceId": "2",
				}
			]
		},
		{
			"roomId": "2",
			"services": [
				{
					"serviceId": "3",
				},
				{
					"serviceId": "4",
				}
			]
		}
	]
	}
*/
