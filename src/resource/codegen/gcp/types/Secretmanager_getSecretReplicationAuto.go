package types

type Secretmanager_getSecretReplicationAuto struct {
	/*
	   The customer-managed encryption configuration of the Secret.
	   If no configuration is provided, Google-managed default
	   encryption is used.
	*/
	CustomerManagedEncryptions []Secretmanager_getSecretReplicationAutoCustomerManagedEncryption `json:"customerManagedEncryptions,omitempty" yaml:"customerManagedEncryptions,omitempty"`
}
