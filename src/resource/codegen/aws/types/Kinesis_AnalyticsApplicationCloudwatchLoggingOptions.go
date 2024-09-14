package types

type Kinesis_AnalyticsApplicationCloudwatchLoggingOptions struct {
	// The ARN of the CloudWatch Log Stream.
	LogStreamArn string `json:"logStreamArn,omitempty" yaml:"logStreamArn,omitempty"`

	// The ARN of the IAM Role used to send application messages.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the Kinesis Analytics Application.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
