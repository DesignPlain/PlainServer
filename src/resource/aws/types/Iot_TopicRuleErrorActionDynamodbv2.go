package types

type Iot_TopicRuleErrorActionDynamodbv2 struct {
	// Configuration block with DynamoDB Table to which the message will be written. Nested arguments below.
	PutItem Iot_TopicRuleErrorActionDynamodbv2PutItem `json:"putItem,omitempty" yaml:"putItem,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch alarm.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
