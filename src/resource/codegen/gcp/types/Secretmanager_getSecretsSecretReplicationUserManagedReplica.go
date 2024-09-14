package types

type Secretmanager_getSecretsSecretReplicationUserManagedReplica struct {
	/*
	   Customer Managed Encryption for the secret.
	   Structure is documented below.
	*/
	CustomerManagedEncryptions []Secretmanager_getSecretsSecretReplicationUserManagedReplicaCustomerManagedEncryption `json:"customerManagedEncryptions,omitempty" yaml:"customerManagedEncryptions,omitempty"`

	// The canonical IDs of the location to replicate data.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
