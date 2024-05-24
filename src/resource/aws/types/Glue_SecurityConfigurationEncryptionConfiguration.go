package types

type Glue_SecurityConfigurationEncryptionConfiguration struct {
	//
	CloudwatchEncryption Glue_SecurityConfigurationEncryptionConfigurationCloudwatchEncryption `json:"cloudwatchEncryption,omitempty" yaml:"cloudwatchEncryption,omitempty"`

	//
	JobBookmarksEncryption Glue_SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption `json:"jobBookmarksEncryption,omitempty" yaml:"jobBookmarksEncryption,omitempty"`

	// A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
	S3Encryption Glue_SecurityConfigurationEncryptionConfigurationS3Encryption `json:"s3Encryption,omitempty" yaml:"s3Encryption,omitempty"`
}
