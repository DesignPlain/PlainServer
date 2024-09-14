package types

type Alb_getLoadBalancerSubnetMapping struct {
	//
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	//
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	//
	Ipv6Address string `json:"ipv6Address,omitempty" yaml:"ipv6Address,omitempty"`

	//
	OutpostId string `json:"outpostId,omitempty" yaml:"outpostId,omitempty"`

	//
	PrivateIpv4Address string `json:"privateIpv4Address,omitempty" yaml:"privateIpv4Address,omitempty"`
}
