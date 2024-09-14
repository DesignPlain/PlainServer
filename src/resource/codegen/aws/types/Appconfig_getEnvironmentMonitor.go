package types

type Appconfig_getEnvironmentMonitor struct {
	// ARN of the Amazon CloudWatch alarm.
	AlarmArn string `json:"alarmArn,omitempty" yaml:"alarmArn,omitempty"`

	// ARN of an IAM role for AWS AppConfig to monitor.
	AlarmRoleArn string `json:"alarmRoleArn,omitempty" yaml:"alarmRoleArn,omitempty"`
}
