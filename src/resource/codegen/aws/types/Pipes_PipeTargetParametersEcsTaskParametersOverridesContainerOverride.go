package types

type Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverride struct {
	// The soft limit (in MiB) of memory to reserve for the container, instead of the default value from the task definition. You must also specify a container name.
	MemoryReservation int `json:"memoryReservation,omitempty" yaml:"memoryReservation,omitempty"`

	// Name of the pipe. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type and amount of a resource to assign to a container, instead of the default value from the task definition. The only supported resource is a GPU. Detailed below.
	ResourceRequirements []Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverrideResourceRequirement `json:"resourceRequirements,omitempty" yaml:"resourceRequirements,omitempty"`

	// List of commands to send to the container that overrides the default command from the Docker image or the task definition. You must also specify a container name.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The number of cpu units reserved for the container, instead of the default value from the task definition. You must also specify a container name.
	Cpu int `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// A list of files containing the environment variables to pass to a container, instead of the value from the container definition. Detailed below.
	EnvironmentFiles []Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironmentFile `json:"environmentFiles,omitempty" yaml:"environmentFiles,omitempty"`

	// The environment variables to send to the container. You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. You must also specify a container name. Detailed below.
	Environments []Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverrideEnvironment `json:"environments,omitempty" yaml:"environments,omitempty"`

	// The hard limit (in MiB) of memory to present to the container, instead of the default value from the task definition. If your container attempts to exceed the memory specified here, the container is killed. You must also specify a container name.
	Memory int `json:"memory,omitempty" yaml:"memory,omitempty"`
}
