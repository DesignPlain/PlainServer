package types

type Eks_getNodeGroupScalingConfig struct {
	// Minimum number of worker nodes.
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// Desired number of worker nodes.
	DesiredSize int `json:"desiredSize,omitempty" yaml:"desiredSize,omitempty"`

	// Maximum number of worker nodes.
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
}
