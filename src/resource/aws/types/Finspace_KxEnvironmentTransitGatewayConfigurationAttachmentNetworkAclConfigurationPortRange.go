package types

type Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfigurationPortRange struct {
	// First port in the range.
	From int `json:"from,omitempty" yaml:"from,omitempty"`

	// Last port in the range.
	To int `json:"to,omitempty" yaml:"to,omitempty"`
}
