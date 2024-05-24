package directconnect

type BgpPeer struct {
	/*
	   The IPv4 CIDR address to use to send traffic to Amazon.
	   Required for IPv4 BGP peers on public virtual interfaces.
	*/
	AmazonAddress string `json:"amazonAddress,omitempty" yaml:"amazonAddress,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `json:"bgpAsn,omitempty" yaml:"bgpAsn,omitempty"`

	// The authentication key for BGP configuration.
	BgpAuthKey string `json:"bgpAuthKey,omitempty" yaml:"bgpAuthKey,omitempty"`

	/*
	   The IPv4 CIDR destination address to which Amazon should send traffic.
	   Required for IPv4 BGP peers on public virtual interfaces.
	*/
	CustomerAddress string `json:"customerAddress,omitempty" yaml:"customerAddress,omitempty"`

	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId string `json:"virtualInterfaceId,omitempty" yaml:"virtualInterfaceId,omitempty"`

	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`
}
