package types

type Secretmanager_getSecretsSecretReplicationAuto struct {
	/*
	   Customer Managed Encryption for the secret.
	   Structure is documented below.
	*/
	CustomerManagedEncryptions []Secretmanager_getSecretsSecretReplicationAutoCustomerManagedEncryption `json:"customerManagedEncryptions,omitempty" yaml:"customerManagedEncryptions,omitempty"`
}
