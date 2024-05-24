package types

type Ecs_getTaskExecutionNetworkConfiguration struct {
	/*
	   Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.

	   For more information, see the [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) documentation.
	*/
	AssignPublicIp bool `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`

	// Security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// Subnets associated with the task or service.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
