package types

type Iot_TopicRuleErrorActionDynamodbv2 struct {
	// The ARN of the IAM role that grants access to the DynamoDB table.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Configuration block with DynamoDB Table to which the message will be written. Nested arguments below.
	PutItem Iot_TopicRuleErrorActionDynamodbv2PutItem `json:"putItem,omitempty" yaml:"putItem,omitempty"`
}
