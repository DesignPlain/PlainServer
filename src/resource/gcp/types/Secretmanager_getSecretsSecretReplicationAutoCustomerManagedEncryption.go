package types

type Secretmanager_getSecretsSecretReplicationAutoCustomerManagedEncryption struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination secret.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
