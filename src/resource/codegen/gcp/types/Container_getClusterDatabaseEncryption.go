package types

type Container_getClusterDatabaseEncryption struct {
	// The key to use to encrypt/decrypt secrets.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// ENCRYPTED or DECRYPTED.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
