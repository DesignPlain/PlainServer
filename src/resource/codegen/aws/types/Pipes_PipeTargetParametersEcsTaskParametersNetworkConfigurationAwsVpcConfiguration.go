package types

type Pipes_PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfiguration struct {
	// Specifies whether the task's elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
	AssignPublicIp string `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`

	//
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
