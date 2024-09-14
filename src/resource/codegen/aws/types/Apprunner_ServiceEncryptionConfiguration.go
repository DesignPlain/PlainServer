package types

type Apprunner_ServiceEncryptionConfiguration struct {
	// ARN of the KMS key used for encryption.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
