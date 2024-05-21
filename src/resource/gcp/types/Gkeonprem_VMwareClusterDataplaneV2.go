package types

type Gkeonprem_VMwareClusterDataplaneV2 struct {
	// Enable advanced networking which requires dataplane_v2_enabled to be set true.
	AdvancedNetworking bool `json:"advancedNetworking,omitempty" yaml:"advancedNetworking,omitempty"`

	// Enables Dataplane V2.
	DataplaneV2Enabled bool `json:"dataplaneV2Enabled,omitempty" yaml:"dataplaneV2Enabled,omitempty"`

	// Enable Dataplane V2 for clusters with Windows nodes.
	WindowsDataplaneV2Enabled bool `json:"windowsDataplaneV2Enabled,omitempty" yaml:"windowsDataplaneV2Enabled,omitempty"`
}
