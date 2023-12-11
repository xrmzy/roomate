package entity

import "time"

type Room struct {
	ID         string    `json:"id"`
	RoomNumber int       `json:"roomNumber"`
	RoomType   string    `json:"roomType"`
	Capacity   int       `json:"capacity"`
	Facility   string    `json:"facility"`
	Price      int64     `json:"price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	IsDeleted  bool      `json:"isDeleted"`
}
