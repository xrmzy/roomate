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

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required, min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	RoleID          int    `json:"roleId" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// func (u User) IsValidRole() bool {
// 	return u.RoleName == "admin" || u.RoleName == "employee" || u.RoleName == "GA"
// }
