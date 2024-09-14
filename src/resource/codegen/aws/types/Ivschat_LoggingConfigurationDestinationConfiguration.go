package types

type Ivschat_LoggingConfigurationDestinationConfiguration struct {
	// An Amazon CloudWatch Logs destination configuration where chat activity will be logged.
	CloudwatchLogs Ivschat_LoggingConfigurationDestinationConfigurationCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// An Amazon Kinesis Data Firehose destination configuration where chat activity will be logged.
	Firehose Ivschat_LoggingConfigurationDestinationConfigurationFirehose `json:"firehose,omitempty" yaml:"firehose,omitempty"`

	// An Amazon S3 destination configuration where chat activity will be logged.
	S3 Ivschat_LoggingConfigurationDestinationConfigurationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
