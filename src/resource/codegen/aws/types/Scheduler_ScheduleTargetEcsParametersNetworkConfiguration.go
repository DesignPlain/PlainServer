package types

type Scheduler_ScheduleTargetEcsParametersNetworkConfiguration struct {
	// Specifies whether the task's elastic network interface receives a public IP address. This attribute is a boolean type, where `true` maps to `ENABLED` and `false` to `DISABLED`. You can specify `true` only when the `launch_type` is set to `FARGATE`.
	AssignPublicIp bool `json:"assignPublicIp,omitempty" yaml:"assignPublicIp,omitempty"`

	// Set of 1 to 5 Security Group ID-s to be associated with the task. These security groups must all be in the same VPC.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// Set of 1 to 16 subnets to be associated with the task. These subnets must all be in the same VPC.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
