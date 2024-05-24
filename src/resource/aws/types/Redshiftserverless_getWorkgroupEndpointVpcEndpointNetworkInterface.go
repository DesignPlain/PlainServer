package types

type Redshiftserverless_getWorkgroupEndpointVpcEndpointNetworkInterface struct {
	// The availability Zone.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// The unique identifier of the network interface.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	// The unique identifier of the subnet.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
