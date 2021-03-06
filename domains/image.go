package domains

//Image represents OS image that will be used for creating VM instances
type Image struct {
	Slug   string `json:"slug"`
	Distro string `json:"distro"`
	OS     string `json:"os"`
}
