package types

type Edgecontainer_NodePoolLocalDiskEncryption struct {
	/*
	   The Cloud KMS CryptoKey e.g. projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} to use for protecting node local disks.
	   If not specified, a Google-managed key will be used instead.
	*/
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	/*
	   (Output)
	   The Cloud KMS CryptoKeyVersion currently in use for protecting node local disks. Only applicable if kmsKey is set.
	*/
	KmsKeyActiveVersion string `json:"kmsKeyActiveVersion,omitempty" yaml:"kmsKeyActiveVersion,omitempty"`

	/*
	   (Output)
	   Availability of the Cloud KMS CryptoKey. If not KEY_AVAILABLE, then nodes may go offline as they cannot access their local data.
	   This can be caused by a lack of permissions to use the key, or if the key is disabled or deleted.
	*/
	KmsKeyState string `json:"kmsKeyState,omitempty" yaml:"kmsKeyState,omitempty"`
}
