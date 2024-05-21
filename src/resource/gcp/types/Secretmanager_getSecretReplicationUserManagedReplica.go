package types

type Secretmanager_getSecretReplicationUserManagedReplica struct {
	// Customer Managed Encryption for the secret.
	CustomerManagedEncryptions []Secretmanager_getSecretReplicationUserManagedReplicaCustomerManagedEncryption `json:"customerManagedEncryptions,omitempty" yaml:"customerManagedEncryptions,omitempty"`

	// The canonical IDs of the location to replicate data. For example: "us-east1".
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
