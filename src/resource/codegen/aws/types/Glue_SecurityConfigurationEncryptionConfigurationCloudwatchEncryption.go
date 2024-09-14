package types

type Glue_SecurityConfigurationEncryptionConfigurationCloudwatchEncryption struct {
	// Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
	CloudwatchEncryptionMode string `json:"cloudwatchEncryptionMode,omitempty" yaml:"cloudwatchEncryptionMode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
