package types

type Ecs_getTaskExecutionPlacementStrategy struct {
	// The field to apply the placement strategy against.
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	/*
	   The type of placement strategy. Valid values are `random`, `spread`, and `binpack`.

	   For more information, see the [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html) documentation.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
