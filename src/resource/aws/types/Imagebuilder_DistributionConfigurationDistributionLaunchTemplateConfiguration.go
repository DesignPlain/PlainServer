package types

type Imagebuilder_DistributionConfigurationDistributionLaunchTemplateConfiguration struct {
	// The account ID that this configuration applies to.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Indicates whether to set the specified Amazon EC2 launch template as the default launch template. Defaults to `true`.
	Default bool `json:"default,omitempty" yaml:"default,omitempty"`

	// The ID of the Amazon EC2 launch template to use.
	LaunchTemplateId string `json:"launchTemplateId,omitempty" yaml:"launchTemplateId,omitempty"`
}
