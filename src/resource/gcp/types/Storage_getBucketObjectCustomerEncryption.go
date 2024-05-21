package types

type Storage_getBucketObjectCustomerEncryption struct {
	// Base64 encoded customer supplied encryption key.
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	// The encryption algorithm. Default: AES256
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty" yaml:"encryptionAlgorithm,omitempty"`
}
