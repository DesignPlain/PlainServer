package types

type Globalaccelerator_AcceleratorIpSet struct {
	// The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
	IpAddresses []string `json:"ipAddresses,omitempty" yaml:"ipAddresses,omitempty"`

	// The type of IP addresses included in this IP set.
	IpFamily string `json:"ipFamily,omitempty" yaml:"ipFamily,omitempty"`
}
