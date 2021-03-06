package domains

//SSHKey represents the actual data for any SSH key in the system
type SSHKey struct {
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
