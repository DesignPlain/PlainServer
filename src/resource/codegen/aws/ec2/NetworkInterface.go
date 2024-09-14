package ec2

import types "libds/aws/types"

type NetworkInterface struct {
	// Description for the network interface.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of private IPs to assign to the ENI in sequential order.
	Ipv6AddressLists []string `json:"ipv6AddressLists,omitempty" yaml:"ipv6AddressLists,omitempty"`

	// List of private IPs to assign to the ENI without regard to order.
	PrivateIps []string `json:"privateIps,omitempty" yaml:"privateIps,omitempty"`

	// List of security group IDs to assign to the ENI.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// Type of network interface to create. Set to `efa` for Elastic Fabric Adapter. Changing `interface_type` will cause the resource to be destroyed and re-created.
	InterfaceType string `json:"interfaceType,omitempty" yaml:"interfaceType,omitempty"`

	// Number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
	Ipv6AddressCount int `json:"ipv6AddressCount,omitempty" yaml:"ipv6AddressCount,omitempty"`

	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. Addresses are assigned without regard to order. You can't use this option if you're specifying `ipv6_address_count`.
	Ipv6Addresses []string `json:"ipv6Addresses,omitempty" yaml:"ipv6Addresses,omitempty"`

	// Configuration block to define the attachment of the ENI. See Attachment below for more details!
	Attachments []types.Ec2_NetworkInterfaceAttachment `json:"attachments,omitempty" yaml:"attachments,omitempty"`

	// Number of IPv4 prefixes that AWS automatically assigns to the network interface.
	Ipv4PrefixCount int `json:"ipv4PrefixCount,omitempty" yaml:"ipv4PrefixCount,omitempty"`

	//
	PrivateIp string `json:"privateIp,omitempty" yaml:"privateIp,omitempty"`

	// Whether `private_ip_list` is allowed and controls the IPs to assign to the ENI and `private_ips` and `private_ips_count` become read-only. Default false.
	PrivateIpListEnabled bool `json:"privateIpListEnabled,omitempty" yaml:"privateIpListEnabled,omitempty"`

	// List of private IPs to assign to the ENI in sequential order. Requires setting `private_ip_list_enabled` to `true`.
	PrivateIpLists []string `json:"privateIpLists,omitempty" yaml:"privateIpLists,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// One or more IPv4 prefixes assigned to the network interface.
	Ipv4Prefixes []string `json:"ipv4Prefixes,omitempty" yaml:"ipv4Prefixes,omitempty"`

	// Whether `ipv6_address_list` is allowed and controls the IPs to assign to the ENI and `ipv6_addresses` and `ipv6_address_count` become read-only. Default false.
	Ipv6AddressListEnabled bool `json:"ipv6AddressListEnabled,omitempty" yaml:"ipv6AddressListEnabled,omitempty"`

	// Number of IPv6 prefixes that AWS automatically assigns to the network interface.
	Ipv6PrefixCount int `json:"ipv6PrefixCount,omitempty" yaml:"ipv6PrefixCount,omitempty"`

	// One or more IPv6 prefixes assigned to the network interface.
	Ipv6Prefixes []string `json:"ipv6Prefixes,omitempty" yaml:"ipv6Prefixes,omitempty"`

	// Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + `private_ips_count`, as a primary private IP will be assiged to an ENI by default.
	PrivateIpsCount int `json:"privateIpsCount,omitempty" yaml:"privateIpsCount,omitempty"`

	// Whether to enable source destination checking for the ENI. Default true.
	SourceDestCheck bool `json:"sourceDestCheck,omitempty" yaml:"sourceDestCheck,omitempty"`

	/*
	   Subnet ID to create the ENI in.

	   The following arguments are optional:
	*/
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
