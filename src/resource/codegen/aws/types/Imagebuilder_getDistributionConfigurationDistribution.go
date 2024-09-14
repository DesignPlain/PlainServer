package types

type Imagebuilder_getDistributionConfigurationDistribution struct {
	// Nested list of Windows faster-launching configurations to use for AMI distribution.
	FastLaunchConfigurations []Imagebuilder_getDistributionConfigurationDistributionFastLaunchConfiguration `json:"fastLaunchConfigurations,omitempty" yaml:"fastLaunchConfigurations,omitempty"`

	// Nested list of launch template configurations.
	LaunchTemplateConfigurations []Imagebuilder_getDistributionConfigurationDistributionLaunchTemplateConfiguration `json:"launchTemplateConfigurations,omitempty" yaml:"launchTemplateConfigurations,omitempty"`

	// Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
	LicenseConfigurationArns []string `json:"licenseConfigurationArns,omitempty" yaml:"licenseConfigurationArns,omitempty"`

	// AWS Region of distribution.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Nested list of AMI distribution configuration.
	AmiDistributionConfigurations []Imagebuilder_getDistributionConfigurationDistributionAmiDistributionConfiguration `json:"amiDistributionConfigurations,omitempty" yaml:"amiDistributionConfigurations,omitempty"`

	// Nested list of container distribution configurations.
	ContainerDistributionConfigurations []Imagebuilder_getDistributionConfigurationDistributionContainerDistributionConfiguration `json:"containerDistributionConfigurations,omitempty" yaml:"containerDistributionConfigurations,omitempty"`
}
