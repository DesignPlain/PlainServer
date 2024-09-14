package types

type Accesscontextmanager_AccessLevelConditionVpcNetworkSourceVpcSubnetwork struct {
	// Required. Network name to be allowed by this Access Level. Networks of foreign organizations requires `compute.network.get` permission to be granted to caller.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// CIDR block IP subnetwork specification. Must be IPv4.
	VpcIpSubnetworks []string `json:"vpcIpSubnetworks,omitempty" yaml:"vpcIpSubnetworks,omitempty"`
}
