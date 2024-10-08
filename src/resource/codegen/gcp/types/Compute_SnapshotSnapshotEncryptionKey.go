package types

type Compute_SnapshotSnapshotEncryptionKey struct {
	// The name of the encryption key that is stored in Google Cloud KMS.
	KmsKeySelfLink string `json:"kmsKeySelfLink,omitempty" yaml:"kmsKeySelfLink,omitempty"`

	/*
	   The service account used for the encryption request for the given KMS key.
	   If absent, the Compute Engine Service Agent service account is used.
	*/
	KmsKeyServiceAccount string `json:"kmsKeyServiceAccount,omitempty" yaml:"kmsKeyServiceAccount,omitempty"`

	/*
	   Specifies a 256-bit customer-supplied encryption key, encoded in
	   RFC 4648 base64 to either encrypt or decrypt this resource.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	RawKey string `json:"rawKey,omitempty" yaml:"rawKey,omitempty"`

	/*
	   (Output)
	   The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	   encryption key that protects this resource.
	*/
	Sha256 string `json:"sha256,omitempty" yaml:"sha256,omitempty"`
}
