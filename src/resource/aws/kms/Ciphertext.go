package kms

type Ciphertext struct {
	// An optional mapping that makes up the encryption context.
	Context map[string]string `json:"context,omitempty" yaml:"context,omitempty"`

	// Globally unique key ID for the customer master key.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext string `json:"plaintext,omitempty" yaml:"plaintext,omitempty"`
}
