package types

type Ecs_getTaskExecutionPlacementConstraint struct {
	// A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can't specify an expression if the constraint type is `distinctInstance`.
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	// The type of constraint. Valid values are `distinctInstance` or `memberOf`. Use `distinctInstance` to ensure that each task in a particular group is running on a different container instance. Use `memberOf` to restrict the selection to a group of valid candidates.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
