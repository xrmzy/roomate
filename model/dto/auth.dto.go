package dto

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
