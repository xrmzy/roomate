package model

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	RoleID    int
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool `json:"isDeleted"`
}

// func (u User) IsValidRole() bool {
// 	return u.RoleName == "admin" || u.RoleName == "employee" || u.RoleName == "GA"
// }
