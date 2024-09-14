package cloudwatch

type EventBus struct {
	// The partner event source that the new event bus will be matched with. Must match `name`.
	EventSourceName string `json:"eventSourceName,omitempty" yaml:"eventSourceName,omitempty"`

	// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on this event bus. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
	KmsKeyIdentifier string `json:"kmsKeyIdentifier,omitempty" yaml:"kmsKeyIdentifier,omitempty"`

	// The name of the new event bus. The names of custom event buses can't contain the / character. To create a partner event bus, ensure the `name` matches the `event_source_name`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
