package cloudfront

type KeyvaluestoreKey struct {
	// Key to put.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Amazon Resource Name (ARN) of the Key Value Store.
	KeyValueStoreArn string `json:"keyValueStoreArn,omitempty" yaml:"keyValueStoreArn,omitempty"`

	// Value to put.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
