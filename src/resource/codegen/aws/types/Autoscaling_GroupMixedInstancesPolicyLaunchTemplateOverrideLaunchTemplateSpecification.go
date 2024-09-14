package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecification struct {
	// ID of the launch template. Conflicts with `launch_template_name`.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`

	// Name of the launch template. Conflicts with `launch_template_id`.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	//
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
