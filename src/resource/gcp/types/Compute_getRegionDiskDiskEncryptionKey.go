package types

type Compute_getRegionDiskDiskEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   Specifies a 256-bit customer-supplied encryption key, encoded in
	   RFC 4648 base64 to either encrypt or decrypt this resource.
	*/
	RawKey string `json:"rawKey,omitempty" yaml:"rawKey,omitempty"`

	/*
	   The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	   encryption key that protects this resource.
	*/
	Sha256 string `json:"sha256,omitempty" yaml:"sha256,omitempty"`
}
