package types

type Container_getClusterNodeConfigTaint struct {
	// Value for taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Effect for taint.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Key for taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
