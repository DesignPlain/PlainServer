package types

type Compute_RegionSecurityPolicyRuleNetworkMatch struct {
	// IPv4 protocol / IPv6 next header (after extension headers). Each element can be an 8-bit unsigned decimal number (e.g. "6"), range (e.g. "253-254"), or one of the following protocol names: "tcp", "udp", "icmp", "esp", "ah", "ipip", or "sctp".
	IpProtocols []string `json:"ipProtocols,omitempty" yaml:"ipProtocols,omitempty"`

	// BGP Autonomous System Number associated with the source IP address.
	SrcAsns []int `json:"srcAsns,omitempty" yaml:"srcAsns,omitempty"`

	// Source IPv4/IPv6 addresses or CIDR prefixes, in standard text format.
	SrcIpRanges []string `json:"srcIpRanges,omitempty" yaml:"srcIpRanges,omitempty"`

	// Source port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").
	SrcPorts []string `json:"srcPorts,omitempty" yaml:"srcPorts,omitempty"`

	// Two-letter ISO 3166-1 alpha-2 country code associated with the source IP address.
	SrcRegionCodes []string `json:"srcRegionCodes,omitempty" yaml:"srcRegionCodes,omitempty"`

	/*
	   User-defined fields. Each element names a defined field and lists the matching values for that field.
	   Structure is documented below.
	*/
	UserDefinedFields []Compute_RegionSecurityPolicyRuleNetworkMatchUserDefinedField `json:"userDefinedFields,omitempty" yaml:"userDefinedFields,omitempty"`

	// Destination IPv4/IPv6 addresses or CIDR prefixes, in standard text format.
	DestIpRanges []string `json:"destIpRanges,omitempty" yaml:"destIpRanges,omitempty"`

	// Destination port numbers for TCP/UDP/SCTP. Each element can be a 16-bit unsigned decimal number (e.g. "80") or range (e.g. "0-1023").
	DestPorts []string `json:"destPorts,omitempty" yaml:"destPorts,omitempty"`
}
