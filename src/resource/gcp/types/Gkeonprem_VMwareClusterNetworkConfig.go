package types

type Gkeonprem_VMwareClusterNetworkConfig struct {
	/*
	   Configuration for control plane V2 mode.
	   Structure is documented below.
	*/
	ControlPlaneV2Config Gkeonprem_VMwareClusterNetworkConfigControlPlaneV2Config `json:"controlPlaneV2Config,omitempty" yaml:"controlPlaneV2Config,omitempty"`

	/*
	   Configuration settings for a DHCP IP configuration.
	   Structure is documented below.
	*/
	DhcpIpConfig Gkeonprem_VMwareClusterNetworkConfigDhcpIpConfig `json:"dhcpIpConfig,omitempty" yaml:"dhcpIpConfig,omitempty"`

	/*
	   Represents common network settings irrespective of the host's IP address.
	   Structure is documented below.
	*/
	HostConfig Gkeonprem_VMwareClusterNetworkConfigHostConfig `json:"hostConfig,omitempty" yaml:"hostConfig,omitempty"`

	/*
	   All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	   Only a single range is supported. This field cannot be changed after creation.
	*/
	PodAddressCidrBlocks []string `json:"podAddressCidrBlocks,omitempty" yaml:"podAddressCidrBlocks,omitempty"`

	/*
	   All services in the cluster are assigned an RFC1918 IPv4 address
	   from these ranges. Only a single range is supported.. This field
	   cannot be changed after creation.
	*/
	ServiceAddressCidrBlocks []string `json:"serviceAddressCidrBlocks,omitempty" yaml:"serviceAddressCidrBlocks,omitempty"`

	/*
	   Configuration settings for a static IP configuration.
	   Structure is documented below.
	*/
	StaticIpConfig Gkeonprem_VMwareClusterNetworkConfigStaticIpConfig `json:"staticIpConfig,omitempty" yaml:"staticIpConfig,omitempty"`

	/*
	   (Output)
	   vcenter_network specifies vCenter network name. Inherited from the admin cluster.
	*/
	VcenterNetwork string `json:"vcenterNetwork,omitempty" yaml:"vcenterNetwork,omitempty"`
}
