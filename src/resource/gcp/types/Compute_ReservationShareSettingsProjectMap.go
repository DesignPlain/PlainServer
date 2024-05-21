package types

type Compute_ReservationShareSettingsProjectMap struct {
	// The identifier for this object. Format specified above.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The project id/number, should be same as the key of this project config in the project map.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
}
