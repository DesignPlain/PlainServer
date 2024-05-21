package types

type Container_getClusterNodePoolDefaultNodeConfigDefault struct {
	// GCFS configuration for this node.
	GcfsConfigs []Container_getClusterNodePoolDefaultNodeConfigDefaultGcfsConfig `json:"gcfsConfigs,omitempty" yaml:"gcfsConfigs,omitempty"`

	// Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT.
	LoggingVariant string `json:"loggingVariant,omitempty" yaml:"loggingVariant,omitempty"`
}
