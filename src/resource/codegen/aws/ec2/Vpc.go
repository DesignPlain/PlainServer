package ec2

type Vpc struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames bool `json:"enableDnsHostnames,omitempty" yaml:"enableDnsHostnames,omitempty"`

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId string `json:"ipv4IpamPoolId,omitempty" yaml:"ipv4IpamPoolId,omitempty"`

	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy string `json:"instanceTenancy,omitempty" yaml:"instanceTenancy,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
	Ipv4NetmaskLength int `json:"ipv4NetmaskLength,omitempty" yaml:"ipv4NetmaskLength,omitempty"`

	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6_netmask_length`.
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6_cidr_block_network_border_group will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" yaml:"ipv6CidrBlockNetworkBorderGroup,omitempty"`

	// IPAM Pool ID for a IPv6 pool. Conflicts with `assign_generated_ipv6_cidr_block`.
	Ipv6IpamPoolId string `json:"ipv6IpamPoolId,omitempty" yaml:"ipv6IpamPoolId,omitempty"`

	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6_ipam_pool_id`
	AssignGeneratedIpv6CidrBlock bool `json:"assignGeneratedIpv6CidrBlock,omitempty" yaml:"assignGeneratedIpv6CidrBlock,omitempty"`

	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport bool `json:"enableDnsSupport,omitempty" yaml:"enableDnsSupport,omitempty"`

	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics bool `json:"enableNetworkAddressUsageMetrics,omitempty" yaml:"enableNetworkAddressUsageMetrics,omitempty"`

	// Netmask length to request from IPAM Pool. Conflicts with `ipv6_cidr_block`. This can be omitted if IPAM pool as a `allocation_default_netmask_length` set. Valid values are from `44` to `60` in increments of 4.
	Ipv6NetmaskLength int `json:"ipv6NetmaskLength,omitempty" yaml:"ipv6NetmaskLength,omitempty"`
}
