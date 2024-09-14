package types

type Athena_DatabaseEncryptionConfiguration struct {
	// Type of key; one of `SSE_S3`, `SSE_KMS`, `CSE_KMS`
	EncryptionOption string `json:"encryptionOption,omitempty" yaml:"encryptionOption,omitempty"`

	// KMS key ARN or ID; required for key types `SSE_KMS` and `CSE_KMS`.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
