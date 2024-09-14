package types

type Ec2_DefaultNetworkAclIngress struct {
	/*
	   The to port to match.

	   The following arguments are optional:
	*/
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// The CIDR block to match. This must be a valid network mask.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// The ICMP type code to be used. Default 0.
	IcmpCode int `json:"icmpCode,omitempty" yaml:"icmpCode,omitempty"`

	// The ICMP type to be used. Default 0.
	IcmpType int `json:"icmpType,omitempty" yaml:"icmpType,omitempty"`

	/*
	   The IPv6 CIDR block.

	   > For more information on ICMP types and codes, see [Internet Control Message Protocol (ICMP) Parameters](https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml).
	*/
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// The protocol to match. If using the -1 'all' protocol, you must specify a from and to port of 0.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The action to take.
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// The from port to match.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// The rule number. Used for ordering.
	RuleNo int `json:"ruleNo,omitempty" yaml:"ruleNo,omitempty"`
}
