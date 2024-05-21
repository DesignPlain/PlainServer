package types

type Container_getClusterNodePoolDefault struct {
	// Subset of NodeConfig message that has defaults.
	NodeConfigDefaults []Container_getClusterNodePoolDefaultNodeConfigDefault `json:"nodeConfigDefaults,omitempty" yaml:"nodeConfigDefaults,omitempty"`
}
