package types

type Container_NodePoolNetworkConfigAdditionalNodeNetworkConfig struct {
	// Name of the subnetwork where the additional interface belongs.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// Name of the VPC where the additional interface belongs.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
