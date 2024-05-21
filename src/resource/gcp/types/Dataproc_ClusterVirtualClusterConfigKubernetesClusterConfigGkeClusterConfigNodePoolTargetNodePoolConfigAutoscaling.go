package types

type Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling struct {
	// The maximum number of nodes in the node pool. Must be >= minNodeCount, and must be > 0.
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	// The minimum number of nodes in the node pool. Must be >= 0 and <= maxNodeCount.
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
}
