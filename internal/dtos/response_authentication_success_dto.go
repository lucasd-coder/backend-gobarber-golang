package dtos

type ResponseUserAuthenticatedSuccessDTO struct {
	Response ResponseProfileDTO `json:"user" `
	Token    string             `json:"token"`
}

func NewResponseUserAuthenticatedSuccessDTO(response ResponseProfileDTO, token string) *ResponseUserAuthenticatedSuccessDTO {
	return &ResponseUserAuthenticatedSuccessDTO{
		Response: response,
		Token:    token,
	}
}
