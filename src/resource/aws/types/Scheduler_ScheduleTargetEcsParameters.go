package types

type Scheduler_ScheduleTargetEcsParameters struct {
	// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. One of: `EC2`, `FARGATE`, `EXTERNAL`.
	LaunchType string `json:"launchType,omitempty" yaml:"launchType,omitempty"`

	// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as `1.1.0`.
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`

	// The number of tasks to create. Ranges from `1` (default) to `10`.
	TaskCount int `json:"taskCount,omitempty" yaml:"taskCount,omitempty"`

	// Up to `6` capacity provider strategies to use for the task. Detailed below.
	CapacityProviderStrategies []Scheduler_ScheduleTargetEcsParametersCapacityProviderStrategy `json:"capacityProviderStrategies,omitempty" yaml:"capacityProviderStrategies,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the task. For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon ECS Developer Guide.
	EnableEcsManagedTags bool `json:"enableEcsManagedTags,omitempty" yaml:"enableEcsManagedTags,omitempty"`

	// A set of up to 5 placement strategies. Detailed below.
	PlacementStrategies []Scheduler_ScheduleTargetEcsParametersPlacementStrategy `json:"placementStrategies,omitempty" yaml:"placementStrategies,omitempty"`

	/*
	   ARN of the task definition to use.

	   The following arguments are optional:
	*/
	TaskDefinitionArn string `json:"taskDefinitionArn,omitempty" yaml:"taskDefinitionArn,omitempty"`

	// Specifies an ECS task group for the task. At most 255 characters.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// A set of up to 10 placement constraints to use for the task. Detailed below.
	PlacementConstraints []Scheduler_ScheduleTargetEcsParametersPlacementConstraint `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`

	// Specifies whether to enable the execute command functionality for the containers in this task.
	EnableExecuteCommand bool `json:"enableExecuteCommand,omitempty" yaml:"enableExecuteCommand,omitempty"`

	// Reference ID to use for the task.
	ReferenceId string `json:"referenceId,omitempty" yaml:"referenceId,omitempty"`

	// The metadata that you apply to the task. Each tag consists of a key and an optional value. For more information, see [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) in the Amazon ECS API Reference.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configures the networking associated with the task. Detailed below.
	NetworkConfiguration Scheduler_ScheduleTargetEcsParametersNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// Specifies whether to propagate the tags from the task definition to the task. One of: `TASK_DEFINITION`.
	PropagateTags string `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`
}
