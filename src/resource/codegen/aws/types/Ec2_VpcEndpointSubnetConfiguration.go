package types

type Ec2_VpcEndpointSubnetConfiguration struct {
	// The IPv6 address to assign to the endpoint network interface in the subnet. You must provide an IPv6 address if the VPC endpoint supports IPv6.
	Ipv6 string `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`

	//
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The IPv4 address to assign to the endpoint network interface in the subnet. You must provide an IPv4 address if the VPC endpoint supports IPv4.
	Ipv4 string `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
}
