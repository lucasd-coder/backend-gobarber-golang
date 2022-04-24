package dtos

type FindAllInMonthFromProviderDTO struct {
	ProviderId string `json:"provider_id"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
}
