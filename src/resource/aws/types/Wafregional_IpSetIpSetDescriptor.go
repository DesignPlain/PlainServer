package types

type Wafregional_IpSetIpSetDescriptor struct {
	// The string like IPV4 or IPV6.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The CIDR notation.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
