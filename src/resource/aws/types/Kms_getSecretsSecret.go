package types

type Kms_getSecretsSecret struct {
	// An optional mapping that makes up the Encryption Context for the secret.
	Context map[string]string `json:"context,omitempty" yaml:"context,omitempty"`

	// The encryption algorithm that will be used to decrypt the ciphertext. This parameter is required only when the ciphertext was encrypted under an asymmetric KMS key. Valid Values: SYMMETRIC_DEFAULT | RSAES_OAEP_SHA_1 | RSAES_OAEP_SHA_256 | SM2PKE
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty" yaml:"encryptionAlgorithm,omitempty"`

	// An optional list of Grant Tokens for the secret.
	GrantTokens []string `json:"grantTokens,omitempty" yaml:"grantTokens,omitempty"`

	/*
	   Specifies the KMS key that AWS KMS uses to decrypt the ciphertext. This parameter is required only when the ciphertext was encrypted under an asymmetric KMS key.

	   For more information on `context` and `grant_tokens` see the [KMS
	   Concepts](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html)
	*/
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// Name to export this secret under in the attributes.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Base64 encoded payload, as returned from a KMS encrypt operation.
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`
}
