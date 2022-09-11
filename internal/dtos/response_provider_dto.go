package dtos

type ResponseProviderDTO struct {
	Hour      int  `json:"hour"`
	Available bool `json:"available"`
}

func NewResponseProviderDTO(hour int, available bool) *ResponseProviderDTO {
	return &ResponseProviderDTO{
		Hour:      hour,
		Available: available,
	}
}
