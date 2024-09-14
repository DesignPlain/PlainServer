package types

type Iot_TopicRuleErrorActionIotEvents struct {
	// The payload that contains a JSON array of records will be sent to IoT Events via a batch call.
	BatchMode bool `json:"batchMode,omitempty" yaml:"batchMode,omitempty"`

	// The name of the AWS IoT Events input.
	InputName string `json:"inputName,omitempty" yaml:"inputName,omitempty"`

	// Use this to ensure that only one input (message) with a given messageId is processed by an AWS IoT Events detector.
	MessageId string `json:"messageId,omitempty" yaml:"messageId,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
