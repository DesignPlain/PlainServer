package types



type Container_getClusterNodePoolNetworkConfig struct {
	// Configuration for node-pool level pod cidr overprovision. If not set, the cluster level setting will be inherited
	PodCidrOverprovisionConfigs []Container_getClusterNodePoolNetworkConfigPodCidrOverprovisionConfig `json:"podCidrOverprovisionConfigs,omitempty" yaml:"podCidrOverprovisionConfigs,omitempty"`

	// The IP address range for pod IPs in this node pool. Only applicable if create_pod_range is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) to pick a specific range to use.
	PodIpv4CidrBlock string `json:"podIpv4CidrBlock,omitempty" yaml:"podIpv4CidrBlock,omitempty"`

	// The ID of the secondary range for pod IPs. If create_pod_range is true, this ID is used for the new range. If create_pod_range is false, uses an existing secondary range with this ID.
	PodRange string `json:"podRange,omitempty" yaml:"podRange,omitempty"`

	// We specify the additional node networks for this node pool using this list. Each node network corresponds to an additional interface
	AdditionalNodeNetworkConfigs []Container_getClusterNodePoolNetworkConfigAdditionalNodeNetworkConfig `json:"additionalNodeNetworkConfigs,omitempty" yaml:"additionalNodeNetworkConfigs,omitempty"`

	// We specify the additional pod networks for this node pool using this list. Each pod network corresponds to an additional alias IP range for the node
	AdditionalPodNetworkConfigs []Container_getClusterNodePoolNetworkConfigAdditionalPodNetworkConfig `json:"additionalPodNetworkConfigs,omitempty" yaml:"additionalPodNetworkConfigs,omitempty"`

	// Whether to create a new range for pod IPs in this node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block if they are not specified.
	CreatePodRange bool `json:"createPodRange,omitempty" yaml:"createPodRange,omitempty"`

	// Whether nodes have internal IP addresses only.
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty" yaml:"enablePrivateNodes,omitempty"`

	// Network bandwidth tier configuration.
	NetworkPerformanceConfigs []Container_getClusterNodePoolNetworkConfigNetworkPerformanceConfig `json:"networkPerformanceConfigs,omitempty" yaml:"networkPerformanceConfigs,omitempty"`
}
