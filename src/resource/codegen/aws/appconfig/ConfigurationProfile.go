package appconfig

import types "libds/aws/types"

type ConfigurationProfile struct {
	// The identifier for an Key Management Service key to encrypt new configuration data versions in the AppConfig hosted configuration store. This attribute is only used for hosted configuration types. The identifier can be an KMS key ID, alias, or the Amazon Resource Name (ARN) of the key ID or alias.
	KmsKeyIdentifier string `json:"kmsKeyIdentifier,omitempty" yaml:"kmsKeyIdentifier,omitempty"`

	// ARN of an IAM role with permission to access the configuration at the specified `location_uri`. A retrieval role ARN is not required for configurations stored in the AWS AppConfig `hosted` configuration store. It is required for all other sources that store your configuration.
	RetrievalRoleArn string `json:"retrievalRoleArn,omitempty" yaml:"retrievalRoleArn,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Type of configurations contained in the profile. Valid values: `AWS.AppConfig.FeatureFlags` and `AWS.Freeform`.  Default: `AWS.Freeform`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
	Validators []types.Appconfig_ConfigurationProfileValidator `json:"validators,omitempty" yaml:"validators,omitempty"`

	// Application ID. Must be between 4 and 7 characters in length.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Description of the configuration profile. Can be at most 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify `hosted`. For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the ARN. For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>`.
	LocationUri string `json:"locationUri,omitempty" yaml:"locationUri,omitempty"`

	// Name for the configuration profile. Must be between 1 and 128 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
