package types

type Workstations_WorkstationConfigEncryptionKey struct {
	// The name of the Google Cloud KMS encryption key.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// The service account to use with the specified KMS key.
	KmsKeyServiceAccount string `json:"kmsKeyServiceAccount,omitempty" yaml:"kmsKeyServiceAccount,omitempty"`
}
