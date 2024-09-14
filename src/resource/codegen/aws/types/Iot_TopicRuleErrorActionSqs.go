package types

type Iot_TopicRuleErrorActionSqs struct {
	// The URL of the Amazon SQS queue.
	QueueUrl string `json:"queueUrl,omitempty" yaml:"queueUrl,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Specifies whether to use Base64 encoding.
	UseBase64 bool `json:"useBase64,omitempty" yaml:"useBase64,omitempty"`
}
