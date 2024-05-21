package types

type Container_NodePoolNodeConfigTaint struct {
	// Effect for taint.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Key for taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value for taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
