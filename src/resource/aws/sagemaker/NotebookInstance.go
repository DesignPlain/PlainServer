package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type NotebookInstance struct {
	// Information on the IMDS configuration of the notebook instance. Conflicts with `instance_metadata_service_configuration`. see details below.
	InstanceMetadataServiceConfiguration types.Sagemaker_NotebookInstanceInstanceMetadataServiceConfiguration `json:"instanceMetadataServiceConfiguration,omitempty" yaml:"instanceMetadataServiceConfiguration,omitempty"`

	// The name of ML compute instance type.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// The name of a lifecycle configuration to associate with the notebook instance.
	LifecycleConfigName string `json:"lifecycleConfigName,omitempty" yaml:"lifecycleConfigName,omitempty"`

	// Whether root access is `Enabled` or `Disabled` for users of the notebook instance. The default value is `Enabled`.
	RootAccess string `json:"rootAccess,omitempty" yaml:"rootAccess,omitempty"`

	// The associated security groups.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The VPC subnet ID.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	/*
	   An array of up to three Git repositories to associate with the notebook instance.
	   These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.
	*/
	AdditionalCodeRepositories []string `json:"additionalCodeRepositories,omitempty" yaml:"additionalCodeRepositories,omitempty"`

	// The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in [AWS CodeCommit](https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository.
	DefaultCodeRepository string `json:"defaultCodeRepository,omitempty" yaml:"defaultCodeRepository,omitempty"`

	// The name of the notebook instance (must be unique).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of Elastic Inference (EI) instance types to associate with this notebook instance. See [Elastic Inference Accelerator](https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html) for more details. Valid values: `ml.eia1.medium`, `ml.eia1.large`, `ml.eia1.xlarge`, `ml.eia2.medium`, `ml.eia2.large`, `ml.eia2.xlarge`.
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// The ARN of the IAM role to be used by the notebook instance which allows SageMaker to call other services on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Set to `Disabled` to disable internet access to notebook. Requires `security_groups` and `subnet_id` to be set. Supported values: `Enabled` (Default) or `Disabled`. If set to `Disabled`, the notebook instance will be able to access resources only in your VPC, and will not be able to connect to Amazon SageMaker training and endpoint services unless your configure a NAT Gateway in your VPC.
	DirectInternetAccess string `json:"directInternetAccess,omitempty" yaml:"directInternetAccess,omitempty"`

	// The platform identifier of the notebook instance runtime environment. This value can be either `notebook-al1-v1`, `notebook-al2-v1`, or  `notebook-al2-v2`, depending on which version of Amazon Linux you require.
	PlatformIdentifier string `json:"platformIdentifier,omitempty" yaml:"platformIdentifier,omitempty"`
}
