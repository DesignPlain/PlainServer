package types

type Cloudrunv2_ServiceTraffic struct {
	// Indicates a string to be part of the URI to exclusively reference this target.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`

	/*
	   The allocation type for this traffic target.
	   Possible values are: `TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST`, `TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies percent of the traffic to this Revision. This defaults to zero if unspecified.
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`

	// Revision to which to send this portion of traffic, if traffic allocation is by revision.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`
}
