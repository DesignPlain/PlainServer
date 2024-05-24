package types

type Pipes_PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfiguration struct {
	// List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	// Specifies whether the task's elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
	AssignPublicIp string `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`

	// List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
}
