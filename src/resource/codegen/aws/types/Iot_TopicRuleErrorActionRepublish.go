package types

type Iot_TopicRuleErrorActionRepublish struct {
	// The Quality of Service (QoS) level to use when republishing messages. Valid values are 0 or 1. The default value is 0.
	Qos int `json:"qos,omitempty" yaml:"qos,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The name of the MQTT topic the message should be republished to.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}
