package types

type Ecs_TaskSetNetworkConfiguration struct {
	// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. Maximum of 5.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The subnets associated with the task or service. Maximum of 16.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	/*
	   Whether to assign a public IP address to the ENI (`FARGATE` launch type only). Valid values are `true` or `false`. Default `false`.

	   For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html).
	*/
	AssignPublicIp bool `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`
}
