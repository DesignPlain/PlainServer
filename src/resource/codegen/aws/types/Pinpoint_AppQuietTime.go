package types

type Pinpoint_AppQuietTime struct {
	// The default end time for quiet time in ISO 8601 format. Required if `start` is set
	End string `json:"end,omitempty" yaml:"end,omitempty"`

	// The default start time for quiet time in ISO 8601 format. Required if `end` is set
	Start string `json:"start,omitempty" yaml:"start,omitempty"`
}
