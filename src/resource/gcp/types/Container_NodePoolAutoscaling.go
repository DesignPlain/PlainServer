package types

type Container_NodePoolAutoscaling struct {
	/*
	   Location policy specifies the algorithm used when
	   scaling-up the node pool. Location policy is supported only in 1.24.1+ clusters.
	   - "BALANCED" - Is a best effort policy that aims to balance the sizes of available zones.
	   - "ANY" - Instructs the cluster autoscaler to prioritize utilization of unused reservations,
	   and reduce preemption risk for Spot VMs.
	*/
	LocationPolicy string `json:"locationPolicy,omitempty" yaml:"locationPolicy,omitempty"`

	/*
	   Maximum number of nodes per zone in the NodePool.
	   Must be >= min_node_count. Cannot be used with total limits.
	*/
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	/*
	   Minimum number of nodes per zone in the NodePool.
	   Must be >=0 and <= `max_node_count`. Cannot be used with total limits.
	*/
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`

	/*
	   Total maximum number of nodes in the NodePool.
	   Must be >= total_min_node_count. Cannot be used with per zone limits.
	   Total size limits are supported only in 1.24.1+ clusters.
	*/
	TotalMaxNodeCount int `json:"totalMaxNodeCount,omitempty" yaml:"totalMaxNodeCount,omitempty"`

	/*
	   Total minimum number of nodes in the NodePool.
	   Must be >=0 and <= `total_max_node_count`. Cannot be used with per zone limits.
	   Total size limits are supported only in 1.24.1+ clusters.
	*/
	TotalMinNodeCount int `json:"totalMinNodeCount,omitempty" yaml:"totalMinNodeCount,omitempty"`
}
