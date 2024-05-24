package types

type Connect_InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig struct {
	// The type of encryption. Valid Values: `KMS`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`

	// The full ARN of the encryption key. Be sure to provide the full ARN of the encryption key, not just the ID.
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}
