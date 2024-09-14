package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainer struct {
	// An object that represents the compute environment architecture for AWS Batch jobs on Fargate.
	RuntimePlatforms []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerRuntimePlatform `json:"runtimePlatforms,omitempty" yaml:"runtimePlatforms,omitempty"`

	// A list of ulimits to set in the container.
	Ulimits []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerUlimit `json:"ulimits,omitempty" yaml:"ulimits,omitempty"`

	// The network configuration for jobs that are running on Fargate resources.
	NetworkConfigurations []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerNetworkConfiguration `json:"networkConfigurations,omitempty" yaml:"networkConfigurations,omitempty"`

	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	Privileged bool `json:"privileged,omitempty" yaml:"privileged,omitempty"`

	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem bool `json:"readonlyRootFilesystem,omitempty" yaml:"readonlyRootFilesystem,omitempty"`

	// The secrets for the container.
	Secrets []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	// The user name to use inside the container.
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// The platform configuration for jobs that are running on Fargate resources. Jobs that are running on EC2 resources must not specify this parameter.
	FargatePlatformConfigurations []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerFargatePlatformConfiguration `json:"fargatePlatformConfigurations,omitempty" yaml:"fargatePlatformConfigurations,omitempty"`

	// Linux-specific modifications that are applied to the container.
	LinuxParameters []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLinuxParameter `json:"linuxParameters,omitempty" yaml:"linuxParameters,omitempty"`

	// The type and amount of resources to assign to a container.
	ResourceRequirements []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerResourceRequirement `json:"resourceRequirements,omitempty" yaml:"resourceRequirements,omitempty"`

	// The mount points for data volumes in your container.
	MountPoints []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerMountPoint `json:"mountPoints,omitempty" yaml:"mountPoints,omitempty"`

	// A list of data volumes used in a job.
	Volumes []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// The command that's passed to the container.
	Commands []string `json:"commands,omitempty" yaml:"commands,omitempty"`

	// The Amazon Resource Name (ARN) of the execution role that AWS Batch can assume. For jobs that run on Fargate resources, you must provide an execution role.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// The log configuration specification for the container.
	LogConfigurations []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerLogConfiguration `json:"logConfigurations,omitempty" yaml:"logConfigurations,omitempty"`

	// The instance type to use for a multi-node parallel job.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that the container can assume for AWS permissions.
	JobRoleArn string `json:"jobRoleArn,omitempty" yaml:"jobRoleArn,omitempty"`

	// The environment variables to pass to a container.
	Environments []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerEnvironment `json:"environments,omitempty" yaml:"environments,omitempty"`

	// The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate.
	EphemeralStorages []Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerEphemeralStorage `json:"ephemeralStorages,omitempty" yaml:"ephemeralStorages,omitempty"`

	// The image used to start a container.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`
}
