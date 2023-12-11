package dto

type GetAllParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}