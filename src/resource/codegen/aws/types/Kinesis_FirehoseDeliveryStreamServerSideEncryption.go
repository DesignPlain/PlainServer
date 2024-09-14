package types

type Kinesis_FirehoseDeliveryStreamServerSideEncryption struct {
	// Whether to enable encryption at rest. Default is `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Amazon Resource Name (ARN) of the encryption key. Required when `key_type` is `CUSTOMER_MANAGED_CMK`.
	KeyArn string `json:"keyArn,omitempty" yaml:"keyArn,omitempty"`

	// Type of encryption key. Default is `AWS_OWNED_CMK`. Valid values are `AWS_OWNED_CMK` and `CUSTOMER_MANAGED_CMK`
	KeyType string `json:"keyType,omitempty" yaml:"keyType,omitempty"`
}
