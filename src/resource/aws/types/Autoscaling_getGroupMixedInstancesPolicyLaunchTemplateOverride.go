package types

type Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverride struct {
	/*
	   List of instance requirements objects.
	   - `accelerator_count - List of objects describing the minimum and maximum number of accelerators for an instance type.
	*/
	InstanceRequirements []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirement `json:"instanceRequirements,omitempty" yaml:"instanceRequirements,omitempty"`

	// Overriding instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// List of overriding launch template specification objects.
	LaunchTemplateSpecifications []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification `json:"launchTemplateSpecifications,omitempty" yaml:"launchTemplateSpecifications,omitempty"`

	// Number of capacity units, which gives the instance type a proportional weight to other instance
	WeightedCapacity string `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`
}
