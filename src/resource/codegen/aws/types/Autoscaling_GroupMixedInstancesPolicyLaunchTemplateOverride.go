package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverride struct {
	// Override the instance launch template specification in the Launch Template.
	LaunchTemplateSpecification Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty" yaml:"launchTemplateSpecification,omitempty"`

	// Number of capacity units, which gives the instance type a proportional weight to other instance types.
	WeightedCapacity string `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`

	// Override the instance type in the Launch Template with instance types that satisfy the requirements.
	InstanceRequirements Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements `json:"instanceRequirements,omitempty" yaml:"instanceRequirements,omitempty"`

	// Override the instance type in the Launch Template.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}
