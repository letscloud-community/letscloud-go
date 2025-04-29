package domains

type Snapshot struct {
	Slug        string   `json:"slug"`
	Size        int      `json:"size"`
	Label       string   `json:"label"`
	OsReference string   `json:"os_reference"`
	Reference   string   `json:"reference"`
	Build       bool     `json:"build"`
	Locations   []string `json:"locations"`
}
