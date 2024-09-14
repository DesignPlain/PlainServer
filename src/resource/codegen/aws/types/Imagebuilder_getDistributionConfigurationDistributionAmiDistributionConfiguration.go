package types

type Imagebuilder_getDistributionConfigurationDistributionAmiDistributionConfiguration struct {
	// Key-value map of tags to apply to distributed AMI.
	AmiTags map[string]string `json:"amiTags,omitempty" yaml:"amiTags,omitempty"`

	// Description of the container distribution configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ARN of Key Management Service (KMS) Key to encrypt AMI.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Nested list of EC2 launch permissions.
	LaunchPermissions []Imagebuilder_getDistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermission `json:"launchPermissions,omitempty" yaml:"launchPermissions,omitempty"`

	// Name of the distribution configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of target AWS Account identifiers.
	TargetAccountIds []string `json:"targetAccountIds,omitempty" yaml:"targetAccountIds,omitempty"`
}
