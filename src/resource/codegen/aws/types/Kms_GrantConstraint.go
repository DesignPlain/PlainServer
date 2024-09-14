package types

type Kms_GrantConstraint struct {
	// A list of key-value pairs that must match the encryption context in subsequent cryptographic operation requests. The grant allows the operation only when the encryption context in the request is the same as the encryption context specified in this constraint. Conflicts with `encryption_context_subset`.
	EncryptionContextEquals map[string]string `json:"encryptionContextEquals,omitempty" yaml:"encryptionContextEquals,omitempty"`

	// A list of key-value pairs that must be included in the encryption context of subsequent cryptographic operation requests. The grant allows the cryptographic operation only when the encryption context in the request includes the key-value pairs specified in this constraint, although it can include additional key-value pairs. Conflicts with `encryption_context_equals`.
	EncryptionContextSubset map[string]string `json:"encryptionContextSubset,omitempty" yaml:"encryptionContextSubset,omitempty"`
}
