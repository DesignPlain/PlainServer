package types

type Waf_IpSetIpSetDescriptor struct {
	// Type of the IP address - `IPV4` or `IPV6`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// An IPv4 or IPv6 address specified via CIDR notationE.g., `192.0.2.44/32` or `1111:0000:0000:0000:0000:0000:0000:0000/64`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
