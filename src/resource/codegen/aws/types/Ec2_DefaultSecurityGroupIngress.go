package types

type Ec2_DefaultSecurityGroupIngress struct {
	// Protocol. If you select a protocol of "-1" (semantically equivalent to `all`, which is not a valid value here), you must specify a `from_port` and `to_port` equal to `0`. If not `icmp`, `tcp`, `udp`, or `-1` use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Whether the security group itself will be added as a source to this egress rule.
	Self bool `json:"self,omitempty" yaml:"self,omitempty"`

	// End range port (or ICMP code if protocol is `icmp`).
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// List of CIDR blocks.
	CidrBlocks []string `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Description of the security group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of IPv6 CIDR blocks.
	Ipv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" yaml:"ipv6CidrBlocks,omitempty"`

	// List of prefix list IDs (for allowing access to VPC endpoints)
	PrefixListIds []string `json:"prefixListIds,omitempty" yaml:"prefixListIds,omitempty"`

	// Start port (or ICMP type number if protocol is `icmp`)
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// List of security groups. A group name can be used relative to the default VPC. Otherwise, group ID.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
}
