package entity

import "time"

type Service struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
