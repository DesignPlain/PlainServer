package types

type Pipes_PipeTargetParametersEcsTaskParametersOverrides struct {
	// The Amazon Resource Name (ARN) of the IAM role that containers in this task can assume. All containers in this task are granted the permissions that are specified in this role.
	TaskRoleArn string `json:"taskRoleArn,omitempty" yaml:"taskRoleArn,omitempty"`

	// One or more container overrides that are sent to a task. Detailed below.
	ContainerOverrides []Pipes_PipeTargetParametersEcsTaskParametersOverridesContainerOverride `json:"containerOverrides,omitempty" yaml:"containerOverrides,omitempty"`

	// The number of cpu units reserved for the container, instead of the default value from the task definition. You must also specify a container name.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// The ephemeral storage setting override for the task.  Detailed below.
	EphemeralStorage Pipes_PipeTargetParametersEcsTaskParametersOverridesEphemeralStorage `json:"ephemeralStorage,omitempty" yaml:"ephemeralStorage,omitempty"`

	// The Amazon Resource Name (ARN) of the task execution IAM role override for the task.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// List of Elastic Inference accelerator overrides for the task. Detailed below.
	InferenceAcceleratorOverrides []Pipes_PipeTargetParametersEcsTaskParametersOverridesInferenceAcceleratorOverride `json:"inferenceAcceleratorOverrides,omitempty" yaml:"inferenceAcceleratorOverrides,omitempty"`

	// The hard limit (in MiB) of memory to present to the container, instead of the default value from the task definition. If your container attempts to exceed the memory specified here, the container is killed. You must also specify a container name.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
