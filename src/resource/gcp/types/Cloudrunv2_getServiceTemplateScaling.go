package types

type Cloudrunv2_getServiceTemplateScaling struct {
	// Maximum number of serving instances that this resource should have.
	MaxInstanceCount int `json:"maxInstanceCount,omitempty" yaml:"maxInstanceCount,omitempty"`

	// Minimum number of serving instances that this resource should have.
	MinInstanceCount int `json:"minInstanceCount,omitempty" yaml:"minInstanceCount,omitempty"`
}
