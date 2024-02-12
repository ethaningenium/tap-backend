package models

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Redirect string `json:"redirect" validate:"required,url"`
}