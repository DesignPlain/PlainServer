package types

type Compute_NodeGroupShareSettings struct {
	/*
	   A map of project id and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	   Structure is documented below.
	*/
	ProjectMaps []Compute_NodeGroupShareSettingsProjectMap `json:"projectMaps,omitempty" yaml:"projectMaps,omitempty"`

	/*
	   Node group sharing type.
	   Possible values are: `ORGANIZATION`, `SPECIFIC_PROJECTS`, `LOCAL`.
	*/
	ShareType string `json:"shareType,omitempty" yaml:"shareType,omitempty"`
}
