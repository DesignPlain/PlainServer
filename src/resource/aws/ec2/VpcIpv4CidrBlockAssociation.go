package ec2

type VpcIpv4CidrBlockAssociation struct {
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId string `json:"ipv4IpamPoolId,omitempty" yaml:"ipv4IpamPoolId,omitempty"`

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4_ipam_pool_id`.
	Ipv4NetmaskLength int `json:"ipv4NetmaskLength,omitempty" yaml:"ipv4NetmaskLength,omitempty"`

	// The ID of the VPC to make the association with.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4_netmask_length`.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`
}
