package domains

//GetProfileResponse represents the response data from profile GET request
type GetProfileResponse struct {
	CommonResponse
	Data Profile `json:"data"`
}

//GetLocationsResponse represents the response data from the locations GET request
type GetLocationsResponse struct {
	CommonResponse
	Data []Location `json:"data"`
}

//GetLocationPlansResponse represents the response data from the location plans GET request
type GetLocationPlansResponse struct {
	CommonResponse
	Data []LocationPlanWrapper `json:"data"`
}

//GetLocationImagesResponse represents the response data from the locations images GET request
type GetLocationImagesResponse struct {
	CommonResponse
	Data []Image `json:"data"`
}

//GetSSHKeysResponse represents the response data from the ssh keys GET request
type GetSSHKeysResponse struct {
	CommonResponse
	Data []SSHKey `json:"data"`
}

//CreateOrGetSSHKeysResponse represents the response data from the locations GET/POST request
type CreateOrGetSSHKeysResponse struct {
	CommonResponse
	Data SSHKey `json:"data"`
}

//GetInstancesResponse represents the response data from the all instances GET request
type GetInstancesResponse struct {
	CommonResponse
	Data []Instance `json:"data"`
}

//GetInstanceResponse represents the response data from the single instance GET request
type GetInstanceResponse struct {
	CommonResponse
	Data Instance `json:"data"`
}

//CommonResponse represents the common data from the any GET request
type CommonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

//SSHKeyDelResponse represents the response data from the ssh key DELETE request
type SSHKeyDelResponse struct {
	Message string `json:"message,omitempty"`
}