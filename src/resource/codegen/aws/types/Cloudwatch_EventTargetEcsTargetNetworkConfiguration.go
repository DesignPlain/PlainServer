package types

type Cloudwatch_EventTargetEcsTargetNetworkConfiguration struct {
	// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The subnets associated with the task or service.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	/*
	   Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Defaults to `false`.

	   For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
	*/
	AssignPublicIp bool `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`
}
