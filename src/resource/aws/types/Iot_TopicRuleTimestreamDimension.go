package types

type Iot_TopicRuleTimestreamDimension struct {
	// The name of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the HTTP header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
