package types

type Ecs_getTaskExecutionOverridesContainerOverrideResourceRequirement struct {
	// The type of resource to assign to a container. Valid values are `GPU` or `InferenceAccelerator`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The value for the specified resource type. If the `GPU` type is used, the value is the number of physical GPUs the Amazon ECS container agent reserves for the container. The number of GPUs that's reserved for all containers in a task can't exceed the number of available GPUs on the container instance that the task is launched on. If the `InferenceAccelerator` type is used, the value matches the `deviceName` for an InferenceAccelerator specified in a task definition.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
