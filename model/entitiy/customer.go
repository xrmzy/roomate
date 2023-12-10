package entity

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	IsDeleted   bool   `json:"isDeleted"`
}
