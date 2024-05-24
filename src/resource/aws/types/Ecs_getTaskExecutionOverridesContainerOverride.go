package types

type Ecs_getTaskExecutionOverridesContainerOverride struct {
	// The number of cpu units reserved for the container, instead of the default value from the task definition.
	Cpu int `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// The environment variables to send to the container. You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. See below.
	Environments []Ecs_getTaskExecutionOverridesContainerOverrideEnvironment `json:"environments,omitempty" yaml:"environments,omitempty"`

	// The hard limit (in MiB) of memory to present to the container, instead of the default value from the task definition. If your container attempts to exceed the memory specified here, the container is killed.
	Memory int `json:"memory,omitempty" yaml:"memory,omitempty"`

	// The soft limit (in MiB) of memory to reserve for the container, instead of the default value from the task definition.
	MemoryReservation int `json:"memoryReservation,omitempty" yaml:"memoryReservation,omitempty"`

	// The name of the container that receives the override. This parameter is required if any override is specified.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type and amount of a resource to assign to a container, instead of the default value from the task definition. The only supported resource is a GPU. See below.
	ResourceRequirements []Ecs_getTaskExecutionOverridesContainerOverrideResourceRequirement `json:"resourceRequirements,omitempty" yaml:"resourceRequirements,omitempty"`

	// The command to send to the container that overrides the default command from the Docker image or the task definition.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`
}
