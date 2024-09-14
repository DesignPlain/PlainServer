package types

type Ecs_CapacityProviderAutoScalingGroupProvider struct {
	// ARN of the associated auto scaling group.
	AutoScalingGroupArn string `json:"autoScalingGroupArn,omitempty" yaml:"autoScalingGroupArn,omitempty"`

	// Enables or disables a graceful shutdown of instances without disturbing workloads. Valid values are `ENABLED` and `DISABLED`. The default value is `ENABLED` when a capacity provider is created.
	ManagedDraining string `json:"managedDraining,omitempty" yaml:"managedDraining,omitempty"`

	// Configuration block defining the parameters of the auto scaling. Detailed below.
	ManagedScaling Ecs_CapacityProviderAutoScalingGroupProviderManagedScaling `json:"managedScaling,omitempty" yaml:"managedScaling,omitempty"`

	// Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are `ENABLED` and `DISABLED`.
	ManagedTerminationProtection string `json:"managedTerminationProtection,omitempty" yaml:"managedTerminationProtection,omitempty"`
}
