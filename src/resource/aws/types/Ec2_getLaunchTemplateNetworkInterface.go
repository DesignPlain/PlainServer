package types

type Ec2_getLaunchTemplateNetworkInterface struct {
	//
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	Ipv4AddressCount int `json:"ipv4AddressCount,omitempty" yaml:"ipv4AddressCount,omitempty"`

	//
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty" yaml:"ipv4Addresses,omitempty"`

	//
	Ipv6Prefixes []string `json:"ipv6Prefixes,omitempty" yaml:"ipv6Prefixes,omitempty"`

	//
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	//
	Ipv4PrefixCount int `json:"ipv4PrefixCount,omitempty" yaml:"ipv4PrefixCount,omitempty"`

	//
	Ipv4Prefixes []string `json:"ipv4Prefixes,omitempty" yaml:"ipv4Prefixes,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	//
	InterfaceType string `json:"interfaceType,omitempty" yaml:"interfaceType,omitempty"`

	//
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses,omitempty"`

	//
	Ipv6PrefixCount int `json:"ipv6PrefixCount,omitempty" yaml:"ipv6PrefixCount,omitempty"`

	//
	NetworkCardIndex int `json:"networkCardIndex,omitempty" yaml:"networkCardIndex,omitempty"`

	//
	AssociatePublicIpAddress bool `json:"associatePublicIpAddress,omitempty" yaml:"associatePublicIpAddress,omitempty"`

	//
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	//
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	//
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	//
	AssociateCarrierIpAddress string `json:"associateCarrierIpAddress,omitempty" yaml:"associateCarrierIpAddress,omitempty"`

	//
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" yaml:"ipv6AddressCount,omitempty"`
}
