package types

type Container_getClusterNodePoolNodeConfigTaint struct {
	// Key for taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value for taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Effect for taint.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
}
