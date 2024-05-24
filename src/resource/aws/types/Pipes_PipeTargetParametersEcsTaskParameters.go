package types

type Pipes_PipeTargetParametersEcsTaskParameters struct {
	// The ARN of the task definition to use if the event target is an Amazon ECS task.
	TaskDefinitionArn string `json:"taskDefinitionArn,omitempty" yaml:"taskDefinitionArn,omitempty"`

	// List of capacity provider strategies to use for the task. If a capacityProviderStrategy is specified, the launchType parameter must be omitted. If no capacityProviderStrategy or launchType is specified, the defaultCapacityProviderStrategy for the cluster is used. Detailed below.
	CapacityProviderStrategies []Pipes_PipeTargetParametersEcsTaskParametersCapacityProviderStrategy `json:"capacityProviderStrategies,omitempty" yaml:"capacityProviderStrategies,omitempty"`

	// Use this structure if the Amazon ECS task uses the awsvpc network mode. This structure specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. This structure is required if LaunchType is FARGATE because the awsvpc mode is required for Fargate tasks. If you specify NetworkConfiguration when the target ECS task does not use the awsvpc network mode, the task fails. Detailed below.
	NetworkConfiguration Pipes_PipeTargetParametersEcsTaskParametersNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// The placement strategy objects to use for the task. You can specify a maximum of five strategy rules per task. Detailed below.
	PlacementStrategies []Pipes_PipeTargetParametersEcsTaskParametersPlacementStrategy `json:"placementStrategies,omitempty" yaml:"placementStrategies,omitempty"`

	// The reference ID to use for the task. Maximum length of 1,024.
	ReferenceId string `json:"referenceId,omitempty" yaml:"referenceId,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The number of tasks to create based on TaskDefinition. The default is 1.
	TaskCount int `json:"taskCount,omitempty" yaml:"taskCount,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the task. Valid values: true, false.
	EnableEcsManagedTags bool `json:"enableEcsManagedTags,omitempty" yaml:"enableEcsManagedTags,omitempty"`

	// Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task. Valid values: true, false.
	EnableExecuteCommand bool `json:"enableExecuteCommand,omitempty" yaml:"enableExecuteCommand,omitempty"`

	// Specifies an Amazon ECS task group for the task. The maximum length is 255 characters.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the TagResource API action. Valid Values: TASK_DEFINITION
	PropagateTags string `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`

	// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The FARGATE value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. Valid Values: EC2, FARGATE, EXTERNAL
	LaunchType string `json:"launchType,omitempty" yaml:"launchType,omitempty"`

	// The overrides that are associated with a task. Detailed below.
	Overrides Pipes_PipeTargetParametersEcsTaskParametersOverrides `json:"overrides,omitempty" yaml:"overrides,omitempty"`

	// An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime). Detailed below.
	PlacementConstraints []Pipes_PipeTargetParametersEcsTaskParametersPlacementConstraint `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`

	// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This structure is used only if LaunchType is FARGATE.
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`
}
