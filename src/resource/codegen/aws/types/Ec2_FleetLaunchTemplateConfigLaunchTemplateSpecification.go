package types

type Ec2_FleetLaunchTemplateConfigLaunchTemplateSpecification struct {
	// The name of the launch template.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	// The launch template version number, `$Latest`, or `$Default.`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// The ID of the launch template.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`
}
