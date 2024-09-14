package ec2

type SecurityGroupRule struct {
	// Whether the security group itself will be added as a source to this ingress rule. Cannot be specified with `cidr_blocks`, `ipv6_cidr_blocks`, or `source_security_group_id`.
	Self bool `json:"self,omitempty" yaml:"self,omitempty"`

	// Security group id to allow access to/from, depending on the `type`. Cannot be specified with `cidr_blocks`, `ipv6_cidr_blocks`, or `self`.
	SourceSecurityGroupId string `json:"sourceSecurityGroupId,omitempty" yaml:"sourceSecurityGroupId,omitempty"`

	// End port (or ICMP code if protocol is "icmp").
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// List of CIDR blocks. Cannot be specified with `source_security_group_id` or `self`.
	CidrBlocks []string `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Security group to apply this rule to.
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`

	// Protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	/*
	   Type of rule being created. Valid options are `ingress` (inbound)
	   or `egress` (outbound).

	   The following arguments are optional:

	   > --Note-- Although `cidr_blocks`, `ipv6_cidr_blocks`, `prefix_list_ids`, and `source_security_group_id` are all marked as optional, you _must_ provide one of them in order to configure the source of the traffic.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Start port (or ICMP type number if protocol is "icmp" or "icmpv6").
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// List of IPv6 CIDR blocks. Cannot be specified with `source_security_group_id` or `self`.
	Ipv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" yaml:"ipv6CidrBlocks,omitempty"`

	// List of Prefix List IDs.
	PrefixListIds []string `json:"prefixListIds,omitempty" yaml:"prefixListIds,omitempty"`
}
