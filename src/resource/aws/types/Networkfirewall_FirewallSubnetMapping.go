package types

type Networkfirewall_FirewallSubnetMapping struct {
	// The subnet's IP address type. Valida values: `"DUALSTACK"`, `"IPV4"`.
	IpAddressType string `json:"ipAddressType,omitempty" yaml:"ipAddressType,omitempty"`

	// The unique identifier for the subnet.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
