package types

type Docdb_ClusterParameterGroupParameter struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod string `json:"applyMethod,omitempty" yaml:"applyMethod,omitempty"`

	// The name of the DocumentDB parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the DocumentDB parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
