package types

type Container_ClusterNodePoolNodeConfigEffectiveTaint struct {
	// Effect for taint. Accepted values are `NO_SCHEDULE`, `PREFER_NO_SCHEDULE`, and `NO_EXECUTE`.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// The default or custom node affinity label key name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value for taint.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
