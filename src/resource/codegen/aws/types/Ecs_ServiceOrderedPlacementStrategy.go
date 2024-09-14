package types

type Ecs_ServiceOrderedPlacementStrategy struct {
	/*
	   For the `spread` placement strategy, valid values are `instanceId` (or `host`,
	   which has the same effect), or any platform or custom attribute that is applied to a container instance.
	   For the `binpack` type, valid values are `memory` and `cpu`. For the `random` type, this attribute is not
	   needed. For more information, see [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html).

	   > --Note:-- for `spread`, `host` and `instanceId` will be normalized, by AWS, to be `instanceId`. This means the statefile will show `instanceId` but your config will differ if you use `host`.
	*/
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// Type of placement strategy. Must be one of: `binpack`, `random`, or `spread`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
