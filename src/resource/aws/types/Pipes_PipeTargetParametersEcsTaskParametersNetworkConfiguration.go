package types

type Pipes_PipeTargetParametersEcsTaskParametersNetworkConfiguration struct {
	// Use this structure to specify the VPC subnets and security groups for the task, and whether a public IP address is to be used. This structure is relevant only for ECS tasks that use the awsvpc network mode. Detailed below.
	AwsVpcConfiguration Pipes_PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfiguration `json:"awsVpcConfiguration,omitempty" yaml:"awsVpcConfiguration,omitempty"`
}
