package types

type Cloudrunv2_getServiceTrafficStatus struct {
	// Specifies percent of the traffic to this Revision.
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`

	// Revision to which this traffic is sent.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	// Indicates the string used in the URI to exclusively reference this target.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`

	// The allocation type for this traffic target.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Displays the target URI.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
