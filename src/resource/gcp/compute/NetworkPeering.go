package compute

type NetworkPeering struct {
	/*
	   The peer network in the peering. The peer network
	   may belong to a different project.
	*/
	PeerNetwork string `json:"peerNetwork,omitempty" yaml:"peerNetwork,omitempty"`

	// Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks. The default value is IPV4_ONLY. Possible values: ["IPV4_ONLY", "IPV4_IPV6"].
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`

	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`

	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp bool `json:"exportSubnetRoutesWithPublicIp,omitempty" yaml:"exportSubnetRoutesWithPublicIp,omitempty"`

	// Whether to import the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`

	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp bool `json:"importSubnetRoutesWithPublicIp,omitempty" yaml:"importSubnetRoutesWithPublicIp,omitempty"`

	// Name of the peering.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The primary network of the peering.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
