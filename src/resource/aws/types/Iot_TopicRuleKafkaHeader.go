package types

type Iot_TopicRuleKafkaHeader struct {
	// The value of the HTTP header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The name of the HTTP header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
