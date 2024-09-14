package vpc

type SecurityGroupIngressRule struct {
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	// The ID of the security group.
	SecurityGroupId string `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`

	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// The source IPv4 CIDR range.
	CidrIpv4 string `json:"cidrIpv4,omitempty" yaml:"cidrIpv4,omitempty"`

	// The source IPv6 CIDR range.
	CidrIpv6 string `json:"cidrIpv6,omitempty" yaml:"cidrIpv6,omitempty"`

	// The security group rule description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	// The ID of the source prefix list.
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`

	// The source security group that is referenced in the rule.
	ReferencedSecurityGroupId string `json:"referencedSecurityGroupId,omitempty" yaml:"referencedSecurityGroupId,omitempty"`
}
