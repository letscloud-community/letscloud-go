package domains

//LocationPlanWrapper wraps all the plans in a given location
type LocationPlanWrapper struct {
	Slug    string `json:"slug"`
	Country string `json:"country"`
	City    string `json:"city"`
	Plans   []Plan `json:"plans"`
}

//Plan represents the pricing plan for a VM instance that will be selected by the client
type Plan struct {
	Slug         string `json:"slug"`
	Core         int    `json:"core"`
	Memory       int    `json:"memory"`
	Disk         int    `json:"disk"`
	Bandwidth    int    `json:"bandwidth"`
	Shortcode    string `json:"shortcode"`
	MonthlyValue string `json:"monthly_value"`
}
