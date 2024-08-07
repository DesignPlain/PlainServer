package types

type Compute_getSnapshotSourceDiskEncryptionKey struct {
	/*
	   The service account used for the encryption request for the given KMS key.
	   If absent, the Compute Engine Service Agent service account is used.
	*/
	KmsKeyServiceAccount string `json:"kmsKeyServiceAccount,omitempty" yaml:"kmsKeyServiceAccount,omitempty"`

	/*
	   Specifies a 256-bit customer-supplied encryption key, encoded in
	   RFC 4648 base64 to either encrypt or decrypt this resource.
	*/
	RawKey string `json:"rawKey,omitempty" yaml:"rawKey,omitempty"`
}
