package types

type Eks_getNodeGroupTaint struct {
	// The effect of the taint.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// The key of the taint.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value of the taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
