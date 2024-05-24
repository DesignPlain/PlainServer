package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplate struct {
	// Override the instance launch template specification in the Launch Template.
	LaunchTemplateSpecification Autoscaling_GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty" yaml:"launchTemplateSpecification,omitempty"`

	// List of nested arguments provides the ability to specify multiple instance  This will override the same parameter in the launch template. For on-demand instances, Auto Scaling considers the order of preference of instance types to launch based on the order specified in the overrides list. Defined below.
	Overrides []Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverride `json:"overrides,omitempty" yaml:"overrides,omitempty"`
}
