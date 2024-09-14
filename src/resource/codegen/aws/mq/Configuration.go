package mq

type Configuration struct {
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engine_type` `RabbitMQ`.
	AuthenticationStrategy string `json:"authenticationStrategy,omitempty" yaml:"authenticationStrategy,omitempty"`

	// Broker configuration in XML format for `ActiveMQ` or [Cuttlefish](https://github.com/Kyorai/cuttlefish) format for `RabbitMQ`. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// Description of the configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType string `json:"engineType,omitempty" yaml:"engineType,omitempty"`

	// Version of the broker engine.
	EngineVersion string `json:"engineVersion,omitempty" yaml:"engineVersion,omitempty"`

	/*
	   Name of the configuration.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
