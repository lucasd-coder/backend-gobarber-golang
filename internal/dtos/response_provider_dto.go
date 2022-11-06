package dtos

type ResponseAllInDayFromProviderDTO struct {
	Hour      int  `json:"hour"`
	Available bool `json:"available"`
}

type ResponseAllInMonthFromProviderDTO struct {
	Day      int  `json:"day"`
	Available bool `json:"available"`
}

func NewResponseAllInDayFromProviderDTO(hour int, available bool) *ResponseAllInDayFromProviderDTO {
	return &ResponseAllInDayFromProviderDTO{
		Hour:      hour,
		Available: available,
	}
}

func NewResponseAllInMonthFromProviderDTO(day int, available bool) *ResponseAllInMonthFromProviderDTO {
	return &ResponseAllInMonthFromProviderDTO{
		Day: day,
		Available: available,
	}
}
