package domains

// Profile represents the profile data of the client/user
type Profile struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Currency    string `json:"currency"`
	Balance     string `json:"balance"`
}
