package dtos

type FindAllInMonthFromProviderDTO struct {
	ProviderID string `json:"provider_id"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
}
