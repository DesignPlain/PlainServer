package ec2

type KeyPair struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name for the key pair. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
	KeyName string `json:"keyName,omitempty" yaml:"keyName,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `key_name`. If neither `key_name` nor `key_name_prefix` is provided, the provider will create a unique key name.
	KeyNamePrefix string `json:"keyNamePrefix,omitempty" yaml:"keyNamePrefix,omitempty"`

	// The public key material.
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`
}
