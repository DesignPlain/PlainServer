package types

type Iot_TopicRuleKafkaHeader struct {
	// The value of the Kafka header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The key of the Kafka header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
