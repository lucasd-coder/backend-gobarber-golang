package dtos

type FindAllInDayFromProviderDTO struct {
	ProviderID string `json:"provider_id"`
	Day        int    `json:"day"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
}
