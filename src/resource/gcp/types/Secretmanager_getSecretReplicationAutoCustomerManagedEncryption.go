package types

type Secretmanager_getSecretReplicationAutoCustomerManagedEncryption struct {
	// The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
