package types

type Gkeonprem_VMwareClusterNetworkConfigStaticIpConfig struct {
	/*
	   Represents the configuration values for static IP allocation to nodes.
	   Structure is documented below.
	*/
	IpBlocks []Gkeonprem_VMwareClusterNetworkConfigStaticIpConfigIpBlock `json:"ipBlocks,omitempty" yaml:"ipBlocks,omitempty"`
}
