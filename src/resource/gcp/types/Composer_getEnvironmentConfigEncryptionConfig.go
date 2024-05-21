package types

type Composer_getEnvironmentConfigEncryptionConfig struct {
	// Optional. Customer-managed Encryption Key available through Google's Key Management Service. Cannot be updated.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
