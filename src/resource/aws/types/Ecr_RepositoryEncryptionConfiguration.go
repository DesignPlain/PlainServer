package types

type Ecr_RepositoryEncryptionConfiguration struct {
	// The encryption type to use for the repository. Valid values are `AES256` or `KMS`. Defaults to `AES256`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// The ARN of the KMS key to use when `encryption_type` is `KMS`. If not specified, uses the default AWS managed key for ECR.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
