package types

type Compute_ReservationShareSettings struct {
	/*
	   A map of project number and project config. This is only valid when shareType's value is SPECIFIC_PROJECTS.
	   Structure is documented below.
	*/
	ProjectMaps []Compute_ReservationShareSettingsProjectMap `json:"projectMaps,omitempty" yaml:"projectMaps,omitempty"`

	/*
	   Type of sharing for this shared-reservation
	   Possible values are: `LOCAL`, `SPECIFIC_PROJECTS`.
	*/
	ShareType string `json:"shareType,omitempty" yaml:"shareType,omitempty"`
}
