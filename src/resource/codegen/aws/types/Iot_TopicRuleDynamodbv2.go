package types

type Iot_TopicRuleDynamodbv2 struct {
	// Configuration block with DynamoDB Table to which the message will be written. Nested arguments below.
	PutItem Iot_TopicRuleDynamodbv2PutItem `json:"putItem,omitempty" yaml:"putItem,omitempty"`

	// The ARN of the IAM role that grants access to the DynamoDB table.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
