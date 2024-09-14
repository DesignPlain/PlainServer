package types

type Iot_TopicRuleErrorActionCloudwatchAlarm struct {
	// The CloudWatch alarm name.
	AlarmName string `json:"alarmName,omitempty" yaml:"alarmName,omitempty"`

	// The IAM role ARN that allows access to the CloudWatch alarm.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The reason for the alarm change.
	StateReason string `json:"stateReason,omitempty" yaml:"stateReason,omitempty"`

	// The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	StateValue string `json:"stateValue,omitempty" yaml:"stateValue,omitempty"`
}
