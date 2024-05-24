package ec2

type NetworkAclRule struct {
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction string `json:"ruleAction,omitempty" yaml:"ruleAction,omitempty"`

	// The to port to match.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress bool `json:"egress,omitempty" yaml:"egress,omitempty"`

	// The from port to match.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocolE.g., -1
	IcmpType int `json:"icmpType,omitempty" yaml:"icmpType,omitempty"`

	// The ID of the network ACL.
	NetworkAclId string `json:"networkAclId,omitempty" yaml:"networkAclId,omitempty"`

	/*
	   ICMP protocol: The ICMP code. Required if specifying ICMP for the protocolE.g., -1

	   > --NOTE:-- If the value of `protocol` is `-1` or `all`, the `from_port` and `to_port` values will be ignored and the rule will apply to all ports.

	   > --NOTE:-- If the value of `icmp_type` is `-1` (which results in a wildcard ICMP type), the `icmp_code` must also be set to `-1` (wildcard ICMP code).

	   > Note: For more information on ICMP types and codes, see here: https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml
	*/
	IcmpCode int `json:"icmpCode,omitempty" yaml:"icmpCode,omitempty"`

	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock string `json:"ipv6CidrBlock,omitempty" yaml:"ipv6CidrBlock,omitempty"`

	// The protocol. A value of -1 means all protocols.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber int `json:"ruleNumber,omitempty" yaml:"ruleNumber,omitempty"`
}
