package types

type Secretmanager_SecretReplicationUserManagedReplica struct {
	/*
	   Customer Managed Encryption for the secret.
	   Structure is documented below.
	*/
	CustomerManagedEncryption Secretmanager_SecretReplicationUserManagedReplicaCustomerManagedEncryption `json:"customerManagedEncryption,omitempty" yaml:"customerManagedEncryption,omitempty"`

	// The canonical IDs of the location to replicate data. For example: "us-east1".
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
