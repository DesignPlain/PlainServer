package types

type Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationIcmpTypeCode struct {
	// ICMP code. A value of `-1` means all codes for the specified ICMP type.
	Code int `json:"code,omitempty" yaml:"code,omitempty"`

	// ICMP type. A value of `-1` means all
	Type int `json:"type,omitempty" yaml:"type,omitempty"`
}
