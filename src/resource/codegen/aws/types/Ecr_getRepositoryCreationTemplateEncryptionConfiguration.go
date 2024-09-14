package types

type Ecr_getRepositoryCreationTemplateEncryptionConfiguration struct {
	// Encryption type to use for any created repositories, either `AES256` or `KMS`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// If `encryption_type` is `KMS`, the ARN of the KMS key used.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
