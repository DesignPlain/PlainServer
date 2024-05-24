package types

type Ecs_CapacityProviderAutoScalingGroupProviderManagedScaling struct {
	// Target utilization for the capacity provider. A number between 1 and 100.
	TargetCapacity int `json:"targetCapacity,omitempty" yaml:"targetCapacity,omitempty"`

	// Period of time, in seconds, after a newly launched Amazon EC2 instance can contribute to CloudWatch metrics for Auto Scaling group. If this parameter is omitted, the default value of 300 seconds is used.
	InstanceWarmupPeriod int `json:"instanceWarmupPeriod,omitempty" yaml:"instanceWarmupPeriod,omitempty"`

	// Maximum step adjustment size. A number between 1 and 10,000.
	MaximumScalingStepSize int `json:"maximumScalingStepSize,omitempty" yaml:"maximumScalingStepSize,omitempty"`

	// Minimum step adjustment size. A number between 1 and 10,000.
	MinimumScalingStepSize int `json:"minimumScalingStepSize,omitempty" yaml:"minimumScalingStepSize,omitempty"`

	// Whether auto scaling is managed by ECS. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
