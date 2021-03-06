package domains

//Location represents the geographical location that any instance will be hosted
type Location struct {
	Slug      string `json:"slug"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Available bool   `json:"available"`
}
