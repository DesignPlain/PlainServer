package types

type Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfiguration struct {
	// Indicates whether to `allow` or `deny` the traffic that matches the rule.
	RuleAction string `json:"ruleAction,omitempty" yaml:"ruleAction,omitempty"`

	// Rule number for the entry. All the network ACL entries are processed in ascending order by rule number.
	RuleNumber int `json:"ruleNumber,omitempty" yaml:"ruleNumber,omitempty"`

	// The IPv4 network range to allow or deny, in CIDR notation. The specified CIDR block is modified to its canonical form. For example, `100.68.0.18/18` will be converted to `100.68.0.0/18`.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// Defines the ICMP protocol that consists of the ICMP type and code. Defined below.
	IcmpTypeCode Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCode `json:"icmpTypeCode,omitempty" yaml:"icmpTypeCode,omitempty"`

	// Range of ports the rule applies to. Defined below.
	PortRange Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRange `json:"portRange,omitempty" yaml:"portRange,omitempty"`

	// Protocol number. A value of `1` means all the protocols.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
