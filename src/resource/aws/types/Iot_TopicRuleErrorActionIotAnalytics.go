package types

type Iot_TopicRuleErrorActionIotAnalytics struct {
	// The payload that contains a JSON array of records will be sent to IoT Analytics via a batch call.
	BatchMode bool `json:"batchMode,omitempty" yaml:"batchMode,omitempty"`

	// Name of AWS IOT Analytics channel.
	ChannelName string `json:"channelName,omitempty" yaml:"channelName,omitempty"`

	// The ARN of the IAM role that grants access.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
