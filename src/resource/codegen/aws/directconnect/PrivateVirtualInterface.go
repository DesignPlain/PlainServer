package directconnect

type PrivateVirtualInterface struct {
	// Indicates whether to enable or disable SiteLink.
	SitelinkEnabled bool `json:"sitelinkEnabled,omitempty" yaml:"sitelinkEnabled,omitempty"`

	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`

	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `json:"bgpAsn,omitempty" yaml:"bgpAsn,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress string `json:"customerAddress,omitempty" yaml:"customerAddress,omitempty"`

	/*
	   The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	   The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	*/
	Mtu int `json:"mtu,omitempty" yaml:"mtu,omitempty"`

	// The name for the virtual interface.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The VLAN ID.
	Vlan int `json:"vlan,omitempty" yaml:"vlan,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress string `json:"amazonAddress,omitempty" yaml:"amazonAddress,omitempty"`

	// The authentication key for BGP configuration.
	BgpAuthKey string `json:"bgpAuthKey,omitempty" yaml:"bgpAuthKey,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId string `json:"dxGatewayId,omitempty" yaml:"dxGatewayId,omitempty"`
}
