package types

type Imagebuilder_getDistributionConfigurationDistributionLaunchTemplateConfiguration struct {
	// The account ID that this configuration applies to.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Whether the specified Amazon EC2 launch template is set as the default launch template.
	Default bool `json:"default,omitempty" yaml:"default,omitempty"`

	// ID of the Amazon EC2 launch template.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`
}
