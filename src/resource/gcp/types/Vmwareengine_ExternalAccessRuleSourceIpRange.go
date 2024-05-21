package types

type Vmwareengine_ExternalAccessRuleSourceIpRange struct {
	// An IP address range in the CIDR format.
	IpAddressRange string `json:"ipAddressRange,omitempty" yaml:"ipAddressRange,omitempty"`

	// A single IP address.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
