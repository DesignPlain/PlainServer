package xray

type EncryptionConfig struct {
	// An AWS KMS customer master key (CMK) ARN.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`

	// The type of encryption. Set to `KMS` to use your own key for encryption. Set to `NONE` for default encryption.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
