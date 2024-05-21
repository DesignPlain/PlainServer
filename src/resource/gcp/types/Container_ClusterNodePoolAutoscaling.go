package types

type Container_ClusterNodePoolAutoscaling struct {
	// Minimum number of nodes per zone in the node pool. Must be >=0 and <= max_node_count. Cannot be used with total limits.
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`

	// Maximum number of all nodes in the node pool. Must be >= total_min_node_count. Cannot be used with per zone limits.
	TotalMaxNodeCount int `json:"totalMaxNodeCount,omitempty" yaml:"totalMaxNodeCount,omitempty"`

	// Minimum number of all nodes in the node pool. Must be >=0 and <= total_max_node_count. Cannot be used with per zone limits.
	TotalMinNodeCount int `json:"totalMinNodeCount,omitempty" yaml:"totalMinNodeCount,omitempty"`

	// Location policy specifies the algorithm used when scaling-up the node pool. "BALANCED" - Is a best effort policy that aims to balance the sizes of available zones. "ANY" - Instructs the cluster autoscaler to prioritize utilization of unused reservations, and reduces preemption risk for Spot VMs.
	LocationPolicy string `json:"locationPolicy,omitempty" yaml:"locationPolicy,omitempty"`

	// Maximum number of nodes per zone in the node pool. Must be >= min_node_count. Cannot be used with total limits.
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`
}
