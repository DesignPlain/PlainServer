package types

type Secretmanager_SecretReplicationAuto struct {
	/*
	   The customer-managed encryption configuration of the Secret.
	   If no configuration is provided, Google-managed default
	   encryption is used.
	   Structure is documented below.
	*/
	CustomerManagedEncryption Secretmanager_SecretReplicationAutoCustomerManagedEncryption `json:"customerManagedEncryption,omitempty" yaml:"customerManagedEncryption,omitempty"`
}
