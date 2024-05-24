package types

type Autoscaling_getGroupMixedInstancesPolicyLaunchTemplate struct {
	// List of overriding launch template specification objects.
	LaunchTemplateSpecifications []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `json:"launchTemplateSpecifications,omitempty" yaml:"launchTemplateSpecifications,omitempty"`

	// List of properties overriding the same properties in the launch template.
	Overrides []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverride `json:"overrides,omitempty" yaml:"overrides,omitempty"`
}
