package types

type Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification struct {
	// ID of the launch template.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`

	// Name of the launch template.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	// Template version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
