package domains

//SSHKeyCreateRequest is used for sending POST Request to create new SSH Key
type SSHKeyCreateRequest struct {
	Title string `json:"title"`
	Key   string `json:"key,omitempty"`
}

//SSHKeyDelRequest is used for sending DELETE Request to delete an existing slug
type SSHKeyDelRequest struct {
	Slug string `json:"slug,omitempty"`
}

//CreateInstanceRequest is used for sending POST Request to create a new instance
type CreateInstanceRequest struct {
	LocationSlug string `json:"location_slug"`
	PlanSlug     string `json:"plan_slug"`
	Hostname     string `json:"hostname"`
	Label        string `json:"label"`
	ImageSlug    string `json:"image_slug"`
	SSHSlug      string `json:"ssh_slug,omitempty"`
	Password     string `json:"password,omitempty" validate:"omitempty,min=8"`
}

//InstanceResetPasswordRequest is used for sending PUT Request to reset the password of an existing instance
type InstanceResetPasswordRequest struct {
	Password string `json:"password,omitempty"`
}
