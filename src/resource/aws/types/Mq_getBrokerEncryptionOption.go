package types

type Mq_getBrokerEncryptionOption struct {
	//
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	//
	UseAwsOwnedKey bool `json:"useAwsOwnedKey,omitempty" yaml:"useAwsOwnedKey,omitempty"`
}
