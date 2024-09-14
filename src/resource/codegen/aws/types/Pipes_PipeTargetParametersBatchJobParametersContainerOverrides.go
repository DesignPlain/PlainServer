package types

type Pipes_PipeTargetParametersBatchJobParametersContainerOverrides struct {
	// List of commands to send to the container that overrides the default command from the Docker image or the task definition. You must also specify a container name.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The environment variables to send to the container. You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. You must also specify a container name. Detailed below.
	Environments []Pipes_PipeTargetParametersBatchJobParametersContainerOverridesEnvironment `json:"environments,omitempty" yaml:"environments,omitempty"`

	// The instance type to use for a multi-node parallel job. This parameter isn't applicable to single-node container jobs or jobs that run on Fargate resources, and shouldn't be provided.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The type and amount of a resource to assign to a container, instead of the default value from the task definition. The only supported resource is a GPU. Detailed below.
	ResourceRequirements []Pipes_PipeTargetParametersBatchJobParametersContainerOverridesResourceRequirement `json:"resourceRequirements,omitempty" yaml:"resourceRequirements,omitempty"`
}
