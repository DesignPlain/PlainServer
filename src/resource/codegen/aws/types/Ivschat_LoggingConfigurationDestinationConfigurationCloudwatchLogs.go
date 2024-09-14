package types

type Ivschat_LoggingConfigurationDestinationConfigurationCloudwatchLogs struct {
	// Name of the Amazon Cloudwatch Logs destination where chat activity will be logged.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`
}
