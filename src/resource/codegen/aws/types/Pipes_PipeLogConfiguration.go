package types

type Pipes_PipeLogConfiguration struct {
	// Amazon CloudWatch Logs logging configuration settings for the pipe. Detailed below.
	CloudwatchLogsLogDestination Pipes_PipeLogConfigurationCloudwatchLogsLogDestination `json:"cloudwatchLogsLogDestination,omitempty" yaml:"cloudwatchLogsLogDestination,omitempty"`

	// Amazon Kinesis Data Firehose logging configuration settings for the pipe. Detailed below.
	FirehoseLogDestination Pipes_PipeLogConfigurationFirehoseLogDestination `json:"firehoseLogDestination,omitempty" yaml:"firehoseLogDestination,omitempty"`

	// String list that specifies whether the execution data (specifically, the `payload`, `awsRequest`, and `awsResponse` fields) is included in the log messages for this pipe. This applies to all log destinations for the pipe. Valid values `ALL`.
	IncludeExecutionDatas []string `json:"includeExecutionDatas,omitempty" yaml:"includeExecutionDatas,omitempty"`

	// The level of logging detail to include. Valid values `OFF`, `ERROR`, `INFO` and `TRACE`.
	Level string `json:"level,omitempty" yaml:"level,omitempty"`

	// Amazon S3 logging configuration settings for the pipe. Detailed below.
	S3LogDestination Pipes_PipeLogConfigurationS3LogDestination `json:"s3LogDestination,omitempty" yaml:"s3LogDestination,omitempty"`
}
