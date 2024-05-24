package types

type Networkmanager_ConnectPeerConfigurationBgpConfiguration struct {
	//
	CoreNetworkAsn int `json:"coreNetworkAsn,omitempty" yaml:"coreNetworkAsn,omitempty"`

	/*
	   The Connect peer address.

	   The following arguments are optional:
	*/
	PeerAddress string `json:"peerAddress,omitempty" yaml:"peerAddress,omitempty"`

	//
	PeerAsn int `json:"peerAsn,omitempty" yaml:"peerAsn,omitempty"`

	// A Connect peer core network address.
	CoreNetworkAddress string `json:"coreNetworkAddress,omitempty" yaml:"coreNetworkAddress,omitempty"`
}
