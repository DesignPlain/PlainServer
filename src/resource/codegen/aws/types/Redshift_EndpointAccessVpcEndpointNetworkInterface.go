package types

type Redshift_EndpointAccessVpcEndpointNetworkInterface struct {
	// The Availability Zone.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The network interface identifier.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	// The subnet identifier.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
