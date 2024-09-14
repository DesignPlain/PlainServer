package types

type Imagebuilder_getDistributionConfigurationDistributionFastLaunchConfigurationLaunchTemplate struct {
	// ID of the Amazon EC2 launch template.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`

	// The name of the launch template to use for faster launching for a Windows AMI.
	LaunchTemplateName string `json:"launchTemplateName,omitempty" yaml:"launchTemplateName,omitempty"`

	// The version of the launch template to use for faster launching for a Windows AMI.
	LaunchTemplateVersion string `json:"launchTemplateVersion,omitempty" yaml:"launchTemplateVersion,omitempty"`
}
