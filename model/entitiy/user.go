package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    int       `json:"roleId"`
	RoleName  string    `json:"roleName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"isDeleted"`
}

// func (u User) IsValidRole() bool {
// 	return u.RoleName == "admin" || u.RoleName == "employee" || u.RoleName == "GA"
// }
