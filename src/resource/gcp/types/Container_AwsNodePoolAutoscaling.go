package types

type Container_AwsNodePoolAutoscaling struct {
	// Maximum number of nodes in the NodePool. Must be >= min_node_count.
	MaxNodeCount int `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`

	// Minimum number of nodes in the NodePool. Must be >= 1 and <= max_node_count.
	MinNodeCount int `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
}
