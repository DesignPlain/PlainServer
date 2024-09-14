package types

type Compute_getReservationShareSetting struct {
	// A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	ProjectMaps []Compute_getReservationShareSettingProjectMap `json:"projectMaps,omitempty" yaml:"projectMaps,omitempty"`

	// Type of sharing for this shared-reservation Possible values: ["LOCAL", "SPECIFIC_PROJECTS"]
	ShareType string `json:"shareType,omitempty" yaml:"shareType,omitempty"`
}
