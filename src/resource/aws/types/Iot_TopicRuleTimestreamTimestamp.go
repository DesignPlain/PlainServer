package types

type Iot_TopicRuleTimestreamTimestamp struct {
	// The precision of the timestamp value that results from the expression described in value. Valid values: `SECONDS`, `MILLISECONDS`, `MICROSECONDS`, `NANOSECONDS`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The value of the HTTP header.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
