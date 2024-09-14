package types

type Storage_BucketObjectCustomerEncryption struct {
	// Encryption algorithm. Default: AES256
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty" yaml:"encryptionAlgorithm,omitempty"`

	// Base64 encoded Customer-Supplied Encryption Key.
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`
}
