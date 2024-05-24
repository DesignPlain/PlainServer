package types

type Imagebuilder_DistributionConfigurationDistribution struct {
	// Set of launch template configuration settings that apply to image distribution. Detailed below.
	LaunchTemplateConfigurations []Imagebuilder_DistributionConfigurationDistributionLaunchTemplateConfiguration `json:"launchTemplateConfigurations,omitempty" yaml:"launchTemplateConfigurations,omitempty"`

	// Set of Amazon Resource Names (ARNs) of License Manager License Configurations.
	LicenseConfigurationArns []string `json:"licenseConfigurationArns,omitempty" yaml:"licenseConfigurationArns,omitempty"`

	/*
	   AWS Region for the distribution.

	   The following arguments are optional:
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Configuration block with Amazon Machine Image (AMI) distribution settings. Detailed below.
	AmiDistributionConfiguration Imagebuilder_DistributionConfigurationDistributionAmiDistributionConfiguration `json:"amiDistributionConfiguration,omitempty" yaml:"amiDistributionConfiguration,omitempty"`

	// Configuration block with container distribution settings. Detailed below.
	ContainerDistributionConfiguration Imagebuilder_DistributionConfigurationDistributionContainerDistributionConfiguration `json:"containerDistributionConfiguration,omitempty" yaml:"containerDistributionConfiguration,omitempty"`

	// Set of Windows faster-launching configurations to use for AMI distribution. Detailed below.
	FastLaunchConfigurations []Imagebuilder_DistributionConfigurationDistributionFastLaunchConfiguration `json:"fastLaunchConfigurations,omitempty" yaml:"fastLaunchConfigurations,omitempty"`
}
