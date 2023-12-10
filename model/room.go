package model

import "time"

type Rooms struct {
	ID         string `json:"id"`
	RoomNumber int    `json:"roomNumber"`
	RoomType   string `json:"roomType"`
	Capacity   int    `json:"capacity"`
	Facility   string `json:"facility"`
	Price      int    `json:"price"`
	Status     string `json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsDeleted  bool `json:"isDeleted"`
}
