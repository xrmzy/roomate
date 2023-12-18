package entity

import "time"

type Room struct {
	Id         string    `json:"id"`
	RoomNumber string    `json:"roomNumber"`
	RoomType   string    `json:"roomType"`
	Capacity   int       `json:"capacity"`
	Facility   string    `json:"facility"`
	Price      int       `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	IsDeleted  bool      `json:"isDeleted"`
}
