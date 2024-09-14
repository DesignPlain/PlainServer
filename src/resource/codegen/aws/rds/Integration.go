package rds

import types "libds/aws/types"

type Integration struct {
	// Set of non-secret key–value pairs that contains additional contextual information about the data. For more information, see the [User Guide](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context). You can only include this parameter if you specify the `kms_key_id` parameter.
	AdditionalEncryptionContext map[string]string `json:"additionalEncryptionContext,omitempty" yaml:"additionalEncryptionContext,omitempty"`

	// Name of the integration.
	IntegrationName string `json:"integrationName,omitempty" yaml:"integrationName,omitempty"`

	// KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, RDS uses a default AWS owned key. If you use the default AWS owned key, you should ignore `kms_key_id` parameter by using `lifecycle` parameter to avoid unintended change after the first creation.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// ARN of the database to use as the source for replication.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   ARN of the Redshift data warehouse to use as the target for replication.

	   The following arguments are optional:
	*/
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`

	//
	Timeouts types.Rds_IntegrationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
