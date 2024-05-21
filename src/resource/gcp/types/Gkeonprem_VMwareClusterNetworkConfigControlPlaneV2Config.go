package types

type Gkeonprem_VMwareClusterNetworkConfigControlPlaneV2Config struct {
	/*
	   Static IP addresses for the control plane nodes.
	   Structure is documented below.
	*/
	ControlPlaneIpBlock Gkeonprem_VMwareClusterNetworkConfigControlPlaneV2ConfigControlPlaneIpBlock `json:"controlPlaneIpBlock,omitempty" yaml:"controlPlaneIpBlock,omitempty"`
}
