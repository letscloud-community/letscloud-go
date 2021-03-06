package domains

//Instance represents the actual virtual machine (VM) instance
type Instance struct {
	Identifier    string      `json:"identifier"`
	Booted        bool        `json:"booted"`
	Built         bool        `json:"built"`
	Locked        bool        `json:"locked"`
	Suspended     bool        `json:"suspended"`
	Memory        int         `json:"memory"`
	TotalDiskSize int         `json:"total_disk_size"`
	CPUS          int         `json:"cpus"`
	Label         string      `json:"label"`
	IPAddresses   []IPAddress `json:"ip_addresses"`
	TemplateLabel string      `json:"template_label"`
	Hostname      string      `json:"hostname"`
	RootPassword  string      `json:"initial_root_password"`
}

//IPAddress represents the address of any given instance
type IPAddress struct {
	Address string `json:"address"`
}
