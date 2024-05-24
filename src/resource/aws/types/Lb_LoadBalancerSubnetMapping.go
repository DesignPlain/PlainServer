package types

type Lb_LoadBalancerSubnetMapping struct {
	// ID of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The allocation ID of the Elastic IP address for an internet-facing load balancer.
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	// The IPv6 address. You associate IPv6 CIDR blocks with your VPC and choose the subnets where you launch both internet-facing and internal Application Load Balancers or Network Load Balancers.
	Ipv6Address string `json:"ipv6Address,omitempty" yaml:"ipv6Address,omitempty"`

	//
	OutpostId string `json:"outpostId,omitempty" yaml:"outpostId,omitempty"`

	// The private IPv4 address for an internal load balancer.
	PrivateIpv4Address string `json:"privateIpv4Address,omitempty" yaml:"privateIpv4Address,omitempty"`
}
