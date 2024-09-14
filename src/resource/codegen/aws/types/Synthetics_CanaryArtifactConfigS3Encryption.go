package types

type Synthetics_CanaryArtifactConfigS3Encryption struct {
	// The encryption method to use for artifacts created by this canary. Valid values are: `SSE_S3` and `SSE_KMS`.
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	// The ARN of the customer-managed KMS key to use, if you specify `SSE_KMS` for `encryption_mode`.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
