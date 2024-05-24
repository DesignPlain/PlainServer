package types

type Networkmanager_ConnectPeerConfiguration struct {
	// The inside IP addresses used for BGP peering. Required when the Connect attachment protocol is `GRE`. See `aws.networkmanager.ConnectAttachment` for details.
	InsideCidrBlocks []string `json:"insideCidrBlocks,omitempty" yaml:"insideCidrBlocks,omitempty"`

	/*
	   The Connect peer address.

	   The following arguments are optional:
	*/
	PeerAddress string `json:"peerAddress,omitempty" yaml:"peerAddress,omitempty"`

	//
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	//
	BgpConfigurations []Networkmanager_ConnectPeerConfigurationBgpConfiguration `json:"bgpConfigurations,omitempty" yaml:"bgpConfigurations,omitempty"`

	// A Connect peer core network address.
	CoreNetworkAddress string `json:"coreNetworkAddress,omitempty" yaml:"coreNetworkAddress,omitempty"`
}
