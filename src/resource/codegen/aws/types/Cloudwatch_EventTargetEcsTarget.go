package types

type Cloudwatch_EventTargetEcsTarget struct {
	// A map of tags to assign to ecs resources.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
	TaskDefinitionArn string `json:"taskDefinitionArn,omitempty" yaml:"taskDefinitionArn,omitempty"`

	// The capacity provider strategy to use for the task. If a `capacity_provider_strategy` specified, the `launch_type` parameter must be omitted. If no `capacity_provider_strategy` or `launch_type` is specified, the default capacity provider strategy for the cluster is used. Can be one or more. See below.
	CapacityProviderStrategies []Cloudwatch_EventTargetEcsTargetCapacityProviderStrategy `json:"capacityProviderStrategies,omitempty" yaml:"capacityProviderStrategies,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the task.
	EnableEcsManagedTags bool `json:"enableEcsManagedTags,omitempty" yaml:"enableEcsManagedTags,omitempty"`

	// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values include: `EC2`, `EXTERNAL`, or `FARGATE`.
	LaunchType string `json:"launchType,omitempty" yaml:"launchType,omitempty"`

	// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if `launch_type` is `FARGATE` because the awsvpc mode is required for Fargate tasks.
	NetworkConfiguration Cloudwatch_EventTargetEcsTargetNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime). See Below.
	PlacementConstraints []Cloudwatch_EventTargetEcsTargetPlacementConstraint `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`

	// Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. The only valid value is: `TASK_DEFINITION`.
	PropagateTags string `json:"propagateTags,omitempty" yaml:"propagateTags,omitempty"`

	// Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task.
	EnableExecuteCommand bool `json:"enableExecuteCommand,omitempty" yaml:"enableExecuteCommand,omitempty"`

	// Specifies an ECS task group for the task. The maximum length is 255 characters.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// An array of placement strategy objects to use for the task. You can specify a maximum of five strategy rules per task.
	OrderedPlacementStrategies []Cloudwatch_EventTargetEcsTargetOrderedPlacementStrategy `json:"orderedPlacementStrategies,omitempty" yaml:"orderedPlacementStrategies,omitempty"`

	// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as `1.1.0`. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see [AWS Fargate Platform Versions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`

	// The number of tasks to create based on the TaskDefinition. Defaults to `1`.
	TaskCount int `json:"taskCount,omitempty" yaml:"taskCount,omitempty"`
}
