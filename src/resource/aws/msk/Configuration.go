package msk

type Configuration struct {
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `json:"kafkaVersions,omitempty" yaml:"kafkaVersions,omitempty"`

	// Name of the configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
	ServerProperties string `json:"serverProperties,omitempty" yaml:"serverProperties,omitempty"`

	// Description of the configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
