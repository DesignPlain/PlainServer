package types

type Iot_TopicRuleHttpHttpHeader struct {
	// The name of the HTTP header.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value of the HTTP header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
