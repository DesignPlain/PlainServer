package types

type Glue_SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption struct {
	// Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
	JobBookmarksEncryptionMode string `json:"jobBookmarksEncryptionMode,omitempty" yaml:"jobBookmarksEncryptionMode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
