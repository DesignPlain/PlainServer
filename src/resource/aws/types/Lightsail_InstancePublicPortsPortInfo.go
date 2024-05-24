package types

type Lightsail_InstancePublicPortsPortInfo struct {
	/*
	   Last port in a range of open ports on an instance.

	   The following arguments are optional:
	*/
	ToPort int `json:"toPort,omitempty" yaml:"toPort,omitempty"`

	// Set of CIDR aliases that define access for a preconfigured range of IP addresses.
	CidrListAliases []string `json:"cidrListAliases,omitempty" yaml:"cidrListAliases,omitempty"`

	// Set of CIDR blocks.
	Cidrs []string `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`

	// First port in a range of open ports on an instance.
	FromPort int `json:"fromPort,omitempty" yaml:"fromPort,omitempty"`

	//
	Ipv6Cidrs []string `json:"ipv6Cidrs,omitempty" yaml:"ipv6Cidrs,omitempty"`

	// IP protocol name. Valid values are `tcp`, `all`, `udp`, and `icmp`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
