package types

type Ec2_LaunchTemplateNetworkInterface struct {
	// Whether the network interface should be destroyed on instance termination.
	DeleteOnTermination string `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Description of the network interface.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// One or more private IPv4 addresses to associate. Conflicts with `ipv4_address_count`
	Ipv4Addresses []string `json:"ipv4Addresses,omitempty" yaml:"ipv4Addresses,omitempty"`

	// The number of IPv6 addresses to assign to a network interface. Conflicts with `ipv6_addresses`
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" yaml:"ipv6AddressCount,omitempty"`

	// One or more IPv6 prefixes to be assigned to the network interface. Conflicts with `ipv6_prefix_count`
	Ipv6Prefixes []string `json:"ipv6Prefixes,omitempty" yaml:"ipv6Prefixes,omitempty"`

	// The integer index of the network interface attachment.
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	// The type of network interface. To create an Elastic Fabric Adapter (EFA), specify `efa`.
	InterfaceType string `json:"interfaceType,omitempty" yaml:"interfaceType,omitempty"`

	// The number of IPv4 prefixes to be automatically assigned to the network interface. Conflicts with `ipv4_prefixes`
	Ipv4PrefixCount int `json:"ipv4PrefixCount,omitempty" yaml:"ipv4PrefixCount,omitempty"`

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Conflicts with `ipv6_address_count`
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses,omitempty"`

	// The primary private IPv4 address.
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	// A list of security group IDs to associate.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	/*
	   Associate a Carrier IP address with `eth0` for a new network interface.
	   Use this option when you launch an instance in a Wavelength Zone and want to associate a Carrier IP address with the network interface.
	   Boolean value, can be left unset.
	*/
	AssociateCarrierIpAddress string `json:"associateCarrierIpAddress,omitempty" yaml:"associateCarrierIpAddress,omitempty"`

	/*
	   Associate a public ip address with the network interface.
	   Boolean value, can be left unset.
	*/
	AssociatePublicIpAddress string `json:"associatePublicIpAddress,omitempty" yaml:"associatePublicIpAddress,omitempty"`

	// One or more IPv4 prefixes to be assigned to the network interface. Conflicts with `ipv4_prefix_count`
	Ipv4Prefixes []string `json:"ipv4Prefixes,omitempty" yaml:"ipv4Prefixes,omitempty"`

	// The number of IPv6 prefixes to be automatically assigned to the network interface. Conflicts with `ipv6_prefixes`
	Ipv6PrefixCount int `json:"ipv6PrefixCount,omitempty" yaml:"ipv6PrefixCount,omitempty"`

	// The index of the network card. Some instance types support multiple network cards. The primary network interface must be assigned to network card index 0. The default is network card index 0.
	NetworkCardIndex int `json:"networkCardIndex,omitempty" yaml:"networkCardIndex,omitempty"`

	// The number of secondary private IPv4 addresses to assign to a network interface. Conflicts with `ipv4_addresses`
	Ipv4AddressCount int `json:"ipv4AddressCount,omitempty" yaml:"ipv4AddressCount,omitempty"`

	// The ID of the network interface to attach.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The VPC Subnet ID to associate.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
