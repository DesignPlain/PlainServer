package types

type Gkeonprem_BareMetalClusterNetworkConfig struct {
	/*
	   Configuration for SR-IOV.
	   Structure is documented below.
	*/
	SrIovConfig Gkeonprem_BareMetalClusterNetworkConfigSrIovConfig `json:"srIovConfig,omitempty" yaml:"srIovConfig,omitempty"`

	/*
	   Enables the use of advanced Anthos networking features, such as Bundled
	   Load Balancing with BGP or the egress NAT gateway.
	   Setting configuration for advanced networking features will automatically
	   set this flag.
	*/
	AdvancedNetworking bool `json:"advancedNetworking,omitempty" yaml:"advancedNetworking,omitempty"`

	/*
	   A nested object resource
	   Structure is documented below.
	*/
	IslandModeCidr Gkeonprem_BareMetalClusterNetworkConfigIslandModeCidr `json:"islandModeCidr,omitempty" yaml:"islandModeCidr,omitempty"`

	/*
	   Configuration for multiple network interfaces.
	   Structure is documented below.
	*/
	MultipleNetworkInterfacesConfig Gkeonprem_BareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig `json:"multipleNetworkInterfacesConfig,omitempty" yaml:"multipleNetworkInterfacesConfig,omitempty"`
}
