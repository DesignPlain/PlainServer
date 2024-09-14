package types

type Neptune_ParameterGroupParameter struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod string `json:"applyMethod,omitempty" yaml:"applyMethod,omitempty"`

	// The name of the Neptune parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the Neptune parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
