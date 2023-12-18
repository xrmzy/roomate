package entity

import "time"

type Role struct {
	Id        string    `json:"id"`
	RoleName  string    `json:"roleName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}
