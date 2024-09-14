package types

type Ecr_getRepositoryEncryptionConfiguration struct {
	// Encryption type to use for the repository, either `AES256` or `KMS`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// If `encryption_type` is `KMS`, the ARN of the KMS key used.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
