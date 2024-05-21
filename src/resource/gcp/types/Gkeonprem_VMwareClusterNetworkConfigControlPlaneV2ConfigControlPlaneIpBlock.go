package types

type Gkeonprem_VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock struct {
	/*
	   The node's network configurations used by the VMware User Cluster.
	   Structure is documented below.
	*/
	Ips []Gkeonprem_VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlockIp `json:"ips,omitempty" yaml:"ips,omitempty"`

	// The netmask used by the VMware User Cluster.
	Netmask string `json:"netmask,omitempty" yaml:"netmask,omitempty"`

	// The network gateway used by the VMware User Cluster.
	Gateway string `json:"gateway,omitempty" yaml:"gateway,omitempty"`
}
