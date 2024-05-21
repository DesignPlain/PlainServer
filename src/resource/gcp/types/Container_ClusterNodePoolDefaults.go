package types

type Container_ClusterNodePoolDefaults struct {
	// Subset of NodeConfig message that has defaults.
	NodeConfigDefaults Container_ClusterNodePoolDefaultsNodeConfigDefaults `json:"nodeConfigDefaults,omitempty" yaml:"nodeConfigDefaults,omitempty"`
}
