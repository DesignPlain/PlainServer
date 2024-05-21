package types

type Secretmanager_SecretReplicationUserManagedReplicaCustomerManagedEncryption struct {
	/*
	   Describes the Cloud KMS encryption key that will be used to protect destination secret.

	   - - -
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
