package models

type ErrorResponse struct {
	ErrDesc string `json:"error_description"`
	Error   string `json:"error"`
}