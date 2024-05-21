package types

type Dataproc_getMetastoreServiceEncryptionConfig struct {
	/*
	   The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
	   Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'
	*/
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
