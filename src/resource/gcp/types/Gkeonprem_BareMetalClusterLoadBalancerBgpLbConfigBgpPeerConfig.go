package types

type Gkeonprem_BareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfig struct {
	/*
	   BGP autonomous system number (ASN) for the network that contains the
	   external peer device.
	*/
	Asn int `json:"asn,omitempty" yaml:"asn,omitempty"`

	/*
	   The IP address of the control plane node that connects to the external
	   peer.
	   If you don't specify any control plane nodes, all control plane nodes
	   can connect to the external peer. If you specify one or more IP addresses,
	   only the nodes specified participate in peering sessions.
	*/
	ControlPlaneNodes []string `json:"controlPlaneNodes,omitempty" yaml:"controlPlaneNodes,omitempty"`

	// The IP address of the external peer device.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
