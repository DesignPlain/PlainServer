package types

type Ecs_getTaskExecutionOverrides struct {
	// Amazon Resource Name (ARN) of the task execution role override for the task.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Elastic Inference accelerator override for the task. See below.
	InferenceAcceleratorOverrides []Ecs_getTaskExecutionOverridesInferenceAcceleratorOverride `json:"inferenceAcceleratorOverrides,omitempty" yaml:"inferenceAcceleratorOverrides,omitempty"`

	// The memory override for the task.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`

	// Amazon Resource Name (ARN) of the role that containers in this task can assume.
	TaskRoleArn string `json:"taskRoleArn,omitempty" yaml:"taskRoleArn,omitempty"`

	// One or more container overrides that are sent to a task. See below.
	ContainerOverrides []Ecs_getTaskExecutionOverridesContainerOverride `json:"containerOverrides,omitempty" yaml:"containerOverrides,omitempty"`

	// The CPU override for the task.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
}
