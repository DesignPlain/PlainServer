package ec2

type VpcIpv6CidrBlockAssociation struct {
	// The IPv6 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv6_netmask_length`. This parameter is required if `ipv6_netmask_length` is not set and the IPAM pool does not have `allocation_default_netmask` set.
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// The ID of an IPv6 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts.
	Ipv6IpamPoolId string `json:"ipv6IpamPoolId,omitempty" yaml:"ipv6IpamPoolId,omitempty"`

	// The netmask length of the IPv6 CIDR you want to allocate to this VPC. Requires specifying a `ipv6_ipam_pool_id`. This parameter is optional if the IPAM pool has `allocation_default_netmask` set, otherwise it or `cidr_block` are required
	Ipv6NetmaskLength int `json:"ipv6NetmaskLength,omitempty" yaml:"ipv6NetmaskLength,omitempty"`

	// The ID of the VPC to make the association with.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
