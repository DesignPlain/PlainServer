package types

type Iot_TopicRuleErrorActionKinesis struct {
	// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The name of the Amazon Kinesis stream.
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`

	// The partition key.
	PartitionKey string `json:"partitionKey,omitempty" yaml:"partitionKey,omitempty"`
}
