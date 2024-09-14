package directconnect

type HostedPublicVirtualInterface struct {
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress string `json:"amazonAddress,omitempty" yaml:"amazonAddress,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `json:"bgpAsn,omitempty" yaml:"bgpAsn,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress string `json:"customerAddress,omitempty" yaml:"customerAddress,omitempty"`

	// The AWS account that will own the new virtual interface.
	OwnerAccountId string `json:"ownerAccountId,omitempty" yaml:"ownerAccountId,omitempty"`

	// The VLAN ID.
	Vlan int `json:"vlan,omitempty" yaml:"vlan,omitempty"`

	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily string `json:"addressFamily,omitempty" yaml:"addressFamily,omitempty"`

	// The authentication key for BGP configuration.
	BgpAuthKey string `json:"bgpAuthKey,omitempty" yaml:"bgpAuthKey,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// The name for the virtual interface.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []string `json:"routeFilterPrefixes,omitempty" yaml:"routeFilterPrefixes,omitempty"`
}
