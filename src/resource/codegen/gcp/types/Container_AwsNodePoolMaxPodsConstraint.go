package types

type Container_AwsNodePoolMaxPodsConstraint struct {
	/*
	   The maximum number of pods to schedule on a single node.

	   - - -
	*/
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`
}
