package types

type Container_getClusterNodePoolManagement struct {
	// Whether the nodes will be automatically repaired. Enabled by default.
	AutoRepair bool `json:"autoRepair,omitempty" yaml:"autoRepair,omitempty"`

	// Whether the nodes will be automatically upgraded. Enabled by default.
	AutoUpgrade bool `json:"autoUpgrade,omitempty" yaml:"autoUpgrade,omitempty"`
}
