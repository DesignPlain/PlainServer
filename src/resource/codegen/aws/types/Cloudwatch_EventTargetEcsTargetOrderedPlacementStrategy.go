package types

type Cloudwatch_EventTargetEcsTargetOrderedPlacementStrategy struct {
	// The field to apply the placement strategy against. For the `spread` placement strategy, valid values are `instanceId` (or `host`, which has the same effect), or any platform or custom attribute that is applied to a container instance, such as `attribute:ecs.availability-zone`. For the `binpack` placement strategy, valid values are `cpu` and `memory`. For the `random` placement strategy, this field is not used. For more information, see [Amazon ECS task placement strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	Field string `json:"field,omitempty" yaml:"field,omitempty"`

	// Type of placement strategy. The only valid values at this time are `binpack`, `random` and `spread`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
