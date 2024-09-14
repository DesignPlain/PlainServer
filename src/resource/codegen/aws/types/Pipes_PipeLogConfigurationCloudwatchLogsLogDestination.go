package types

type Pipes_PipeLogConfigurationCloudwatchLogsLogDestination struct {
	// Amazon Web Services Resource Name (ARN) for the CloudWatch log group to which EventBridge sends the log records.
	LogGroupArn string `json:"logGroupArn,omitempty" yaml:"logGroupArn,omitempty"`
}
