package types

type Container_NodePoolNetworkConfigAdditionalPodNetworkConfig struct {
	// Name of the subnetwork where the additional pod network belongs.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// The maximum number of pods per node which use this pod network.
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`

	// The name of the secondary range on the subnet which provides IP address for this pod range.
	SecondaryPodRange string `json:"secondaryPodRange,omitempty" yaml:"secondaryPodRange,omitempty"`
}
