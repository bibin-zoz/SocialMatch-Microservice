package models

type UserVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}
