package dtos

type FindAllInDayFromProviderDTO struct {
	ProviderId string `json:"provider_id"`
	Day        int    `json:"day"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
}
