package types

type Iot_TopicRuleErrorActionCloudwatchLogs struct {
	// The payload that contains a JSON array of records will be sent to CloudWatch via a batch call.
	BatchMode bool `json:"batchMode,omitempty" yaml:"batchMode,omitempty"`

	// The CloudWatch log group name.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch alarm.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
