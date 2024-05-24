package types

type Neptune_ClusterParameterGroupParameter struct {
	// The name of the neptune parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the neptune parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod string `json:"applyMethod,omitempty" yaml:"applyMethod,omitempty"`
}
