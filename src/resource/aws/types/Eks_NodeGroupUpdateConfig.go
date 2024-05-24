package types

type Eks_NodeGroupUpdateConfig struct {
	// Desired max number of unavailable worker nodes during node group update.
	MaxUnavailable int `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`

	// Desired max percentage of unavailable worker nodes during node group update.
	MaxUnavailablePercentage int `json:"maxUnavailablePercentage,omitempty" yaml:"maxUnavailablePercentage,omitempty"`
}
