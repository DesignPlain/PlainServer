package types

type Bigtable_InstanceClusterAutoscalingConfig struct {
	// The target CPU utilization for autoscaling, in percentage. Must be between 10 and 80.
	CpuTarget int `json:"cpuTarget,omitempty" yaml:"cpuTarget,omitempty"`

	// The maximum number of nodes for autoscaling.
	MaxNodes int `json:"maxNodes,omitempty" yaml:"maxNodes,omitempty"`

	// The minimum number of nodes for autoscaling.
	MinNodes int `json:"minNodes,omitempty" yaml:"minNodes,omitempty"`

	/*
	   The target storage utilization for autoscaling, in GB, for each node in a cluster. This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD cluster and between 8192 (8TiB) and 16384 (16 TiB) for an HDD cluster. If not set, whatever is already set for the cluster will not change, or if the cluster is just being created, it will use the default value of 2560 for SSD clusters and 8192 for HDD clusters.

	   !> --Warning--: Only one of `autoscaling_config` or `num_nodes` should be set for a cluster. If both are set, `num_nodes` is ignored. If none is set, autoscaling will be disabled and sized to the current node count.
	*/
	StorageTarget int `json:"storageTarget,omitempty" yaml:"storageTarget,omitempty"`
}
