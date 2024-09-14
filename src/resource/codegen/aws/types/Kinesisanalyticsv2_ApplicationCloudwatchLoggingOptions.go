package types

type Kinesisanalyticsv2_ApplicationCloudwatchLoggingOptions struct {
	//
	CloudwatchLoggingOptionId string `json:"cloudwatchLoggingOptionId,omitempty" yaml:"cloudwatchLoggingOptionId,omitempty"`

	// The ARN of the CloudWatch log stream to receive application messages.
	LogStreamArn string `json:"logStreamArn,omitempty" yaml:"logStreamArn,omitempty"`
}
