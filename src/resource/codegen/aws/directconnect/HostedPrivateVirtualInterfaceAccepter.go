package directconnect

type HostedPrivateVirtualInterfaceAccepter struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId string `json:"dxGatewayId,omitempty" yaml:"dxGatewayId,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `json:"virtualInterfaceId,omitempty" yaml:"virtualInterfaceId,omitempty"`

	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`
}
