package types

type Container_getClusterNodePoolNetworkConfigAdditionalNodeNetworkConfig struct {
	// Name of the VPC where the additional interface belongs.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// Name of the subnetwork where the additional interface belongs.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`
}
