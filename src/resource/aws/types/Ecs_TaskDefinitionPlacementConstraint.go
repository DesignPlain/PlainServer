package types

type Ecs_TaskDefinitionPlacementConstraint struct {
	// Cluster Query Language expression to apply to the constraint. For more information, see [Cluster Query Language in the Amazon EC2 Container Service Developer Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// Type of constraint. Use `memberOf` to restrict selection to a group of valid candidates. Note that `distinctInstance` is not supported in task definitions.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
