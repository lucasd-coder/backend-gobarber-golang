package dtos

type ForgotPasswordEmail struct {
	Email string `json:"email" binding:"required,email"`
}
