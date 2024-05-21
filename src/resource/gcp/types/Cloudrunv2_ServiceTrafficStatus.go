package types

type Cloudrunv2_ServiceTrafficStatus struct {
	// Specifies percent of the traffic to this Revision. This defaults to zero if unspecified.
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`

	// The unique name for the revision. If this field is omitted, it will be automatically generated based on the Service name.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	// Indicates a string to be part of the URI to exclusively reference this target.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`

	/*
	   The allocation type for this traffic target.
	   Possible values are: `TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST`, `TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   (Output)
	   Displays the target URI.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
