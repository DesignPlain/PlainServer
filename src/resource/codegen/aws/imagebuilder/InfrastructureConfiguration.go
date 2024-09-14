package imagebuilder

import types "libds/aws/types"

type InfrastructureConfiguration struct {
	// Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
	InstanceMetadataOptions types.Imagebuilder_InfrastructureConfigurationInstanceMetadataOptions `json:"instanceMetadataOptions,omitempty" yaml:"instanceMetadataOptions,omitempty"`

	// Name of EC2 Key Pair.
	KeyPair string `json:"keyPair,omitempty" yaml:"keyPair,omitempty"`

	// Configuration block with logging settings. Detailed below.
	Logging types.Imagebuilder_InfrastructureConfigurationLogging `json:"logging,omitempty" yaml:"logging,omitempty"`

	/*
	   Name for the configuration.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of EC2 Security Group identifiers.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
	TerminateInstanceOnFailure bool `json:"terminateInstanceOnFailure,omitempty" yaml:"terminateInstanceOnFailure,omitempty"`

	// Key-value map of resource tags to assign to the configuration. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description for the configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of IAM Instance Profile.
	InstanceProfileName string `json:"instanceProfileName,omitempty" yaml:"instanceProfileName,omitempty"`

	// Set of EC2 Instance Types.
	InstanceTypes []string `json:"instanceTypes,omitempty" yaml:"instanceTypes,omitempty"`

	// Key-value map of resource tags to assign to infrastructure created by the configuration.
	ResourceTags map[string]string `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// Amazon Resource Name (ARN) of SNS Topic.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`

	// EC2 Subnet identifier. Also requires `security_group_ids` argument.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
