package types

type Bigquery_TableEncryptionConfiguration struct {
	/*
	   The self link or full name of a key which should be used to
	   encrypt this table.  Note that the default bigquery service account will need to have
	   encrypt/decrypt permissions on this key - you may want to see the
	   `gcp.bigquery.getDefaultServiceAccount` datasource and the
	   `gcp.kms.CryptoKeyIAMBinding` resource.
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	// The self link or full name of the kms key version used to encrypt this table.
	KmsKeyVersion string `json:"kmsKeyVersion,omitempty" yaml:"kmsKeyVersion,omitempty"`
}
