package types

type Iot_TopicRuleCloudwatchLog struct {
	// The CloudWatch log group name.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch alarm.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
