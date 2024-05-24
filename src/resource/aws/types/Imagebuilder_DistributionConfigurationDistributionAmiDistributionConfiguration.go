package types

type Imagebuilder_DistributionConfigurationDistributionAmiDistributionConfiguration struct {
	// Name to apply to the distributed AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of AWS Account identifiers to distribute the AMI.
	TargetAccountIds []string `json:"targetAccountIds,omitempty" yaml:"targetAccountIds,omitempty"`

	// Key-value map of tags to apply to the distributed AMI.
	AmiTags map[string]string `json:"amiTags,omitempty" yaml:"amiTags,omitempty"`

	// Description to apply to the distributed AMI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key to encrypt the distributed AMI.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Configuration block of EC2 launch permissions to apply to the distributed AMI. Detailed below.
	LaunchPermission Imagebuilder_DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermission `json:"launchPermission,omitempty" yaml:"launchPermission,omitempty"`
}
