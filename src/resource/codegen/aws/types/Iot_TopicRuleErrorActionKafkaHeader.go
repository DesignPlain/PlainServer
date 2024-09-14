package types

type Iot_TopicRuleErrorActionKafkaHeader struct {
	// The key of the Kafka header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value of the Kafka header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
