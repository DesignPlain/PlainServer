package ec2

type NatGateway struct {
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivity_type` of `public`.
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType string `json:"connectivityType,omitempty" yaml:"connectivityType,omitempty"`

	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp string `json:"privateIp,omitempty" yaml:"privateIp,omitempty"`

	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds []string `json:"secondaryAllocationIds,omitempty" yaml:"secondaryAllocationIds,omitempty"`

	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount int `json:"secondaryPrivateIpAddressCount,omitempty" yaml:"secondaryPrivateIpAddressCount,omitempty"`

	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses []string `json:"secondaryPrivateIpAddresses,omitempty" yaml:"secondaryPrivateIpAddresses,omitempty"`
}
