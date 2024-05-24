package types

type Glue_SecurityConfigurationEncryptionConfigurationS3Encryption struct {
	// Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
	S3EncryptionMode string `json:"s3EncryptionMode,omitempty" yaml:"s3EncryptionMode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
