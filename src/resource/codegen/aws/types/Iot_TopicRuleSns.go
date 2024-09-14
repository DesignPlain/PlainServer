package types

type Iot_TopicRuleSns struct {
	// The message format of the message to publish. Accepted values are "JSON" and "RAW".
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the SNS topic.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`
}
