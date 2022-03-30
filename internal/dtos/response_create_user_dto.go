package dtos

type ResponseCreateUserDTO struct {
	Name  string `json:"name" `
	Email string `json:"email"`
}

func NewResponseCreateUserDTO(name string, email string) *ResponseCreateUserDTO {
	return &ResponseCreateUserDTO{
		Name:  name,
		Email: email,
	}
}
