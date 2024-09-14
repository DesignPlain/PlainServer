package types

type Gkeonprem_VMwareClusterNetworkConfigDhcpIpConfig struct {
	/*
	   enabled is a flag to mark if DHCP IP allocation is
	   used for VMware user clusters.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
