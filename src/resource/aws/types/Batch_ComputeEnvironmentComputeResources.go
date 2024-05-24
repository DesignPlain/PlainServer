package types

type Batch_ComputeEnvironmentComputeResources struct {
	// Provides information used to select Amazon Machine Images (AMIs) for EC2 instances in the compute environment. If Ec2Configuration isn't specified, the default is ECS_AL2. This parameter isn't applicable to jobs that are running on Fargate resources, and shouldn't be specified.
	Ec2Configurations []Batch_ComputeEnvironmentComputeResourcesEc2Configuration `json:"ec2Configurations,omitempty" yaml:"ec2Configurations,omitempty"`

	// The Amazon Machine Image (AMI) ID used for instances launched in the compute environment. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified. (Deprecated, use `ec2_configuration` `image_id_override` instead)
	ImageId string `json:"imageId,omitempty" yaml:"imageId,omitempty"`

	// Key-value pair tags to be applied to resources that are launched in the compute environment. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The desired number of EC2 vCPUS in the compute environment. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	DesiredVcpus int `json:"desiredVcpus,omitempty" yaml:"desiredVcpus,omitempty"`

	// The minimum number of EC2 vCPUs that an environment should maintain. For `EC2` or `SPOT` compute environments, if the parameter is not explicitly defined, a `0` default value will be set. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	MinVcpus int `json:"minVcpus,omitempty" yaml:"minVcpus,omitempty"`

	// The Amazon EC2 placement group to associate with your compute resources.
	PlacementGroup string `json:"placementGroup,omitempty" yaml:"placementGroup,omitempty"`

	// A list of VPC subnets into which the compute resources are launched.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	// The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. For valid values, refer to the [AWS documentation](https://docs.aws.amazon.com/batch/latest/APIReference/API_ComputeResource.html#Batch-Type-ComputeResource-allocationStrategy). Defaults to `BEST_FIT`. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	AllocationStrategy string `json:"allocationStrategy,omitempty" yaml:"allocationStrategy,omitempty"`

	// The launch template to use for your compute resources. See details below. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	LaunchTemplate Batch_ComputeEnvironmentComputeResourcesLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	SpotIamFleetRole string `json:"spotIamFleetRole,omitempty" yaml:"spotIamFleetRole,omitempty"`

	// The type of compute environment. Valid items are `EC2`, `SPOT`, `FARGATE` or `FARGATE_SPOT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A list of instance types that may be launched. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	InstanceTypes []string `json:"instanceTypes,omitempty" yaml:"instanceTypes,omitempty"`

	// The EC2 key pair that is used for instances launched in the compute environment. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	Ec2KeyPair string `json:"ec2KeyPair,omitempty" yaml:"ec2KeyPair,omitempty"`

	// The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	InstanceRole string `json:"instanceRole,omitempty" yaml:"instanceRole,omitempty"`

	// The maximum number of EC2 vCPUs that an environment can reach.
	MaxVcpus int `json:"maxVcpus,omitempty" yaml:"maxVcpus,omitempty"`

	// A list of EC2 security group that are associated with instances launched in the compute environment. This parameter is required for Fargate compute environments.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Integer of maximum percentage that a Spot Instance price can be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20%!((MISSING)`20`), then the Spot price must be below 20%!o(MISSING)f the current On-Demand price for that EC2 instance. If you leave this field empty, the default value is 100%!o(MISSING)f the On-Demand price. This parameter isn't applicable to jobs running on Fargate resources, and shouldn't be specified.
	BidPercentage int `json:"bidPercentage,omitempty" yaml:"bidPercentage,omitempty"`
}
