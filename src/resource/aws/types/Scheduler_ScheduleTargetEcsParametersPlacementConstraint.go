package types

type Scheduler_ScheduleTargetEcsParametersPlacementConstraint struct {
	// A cluster query language expression to apply to the constraint. You cannot specify an expression if the constraint type is `distinctInstance`. For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the Amazon ECS Developer Guide.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// The type of constraint. One of: `distinctInstance`, `memberOf`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
