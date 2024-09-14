package types

type Cloudwatch_EventTargetEcsTargetPlacementConstraint struct {
	// Cluster Query Language expression to apply to the constraint. Does not need to be specified for the `distinctInstance` type. For more information, see [Cluster Query Language in the Amazon EC2 Container Service Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// Type of constraint. The only valid values at this time are `memberOf` and `distinctInstance`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
