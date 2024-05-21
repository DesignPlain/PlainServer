package types

type Compute_getInstanceGroupManagerAllInstancesConfig struct {
	// The label key-value pairs that you want to patch onto the instance,
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The metadata key-value pairs that you want to patch onto the instance. For more information, see Project and instance metadata,
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
