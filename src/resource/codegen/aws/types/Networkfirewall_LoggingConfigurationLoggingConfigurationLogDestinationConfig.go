package types

type Networkfirewall_LoggingConfigurationLoggingConfigurationLogDestinationConfig struct {
	// The type of log to send. Valid values: `ALERT` or `FLOW` or `TLS`. Alert logs report traffic that matches a `StatefulRule` with an action setting that sends a log message. Flow logs are standard network traffic flow logs.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`

	/*
	   A map describing the logging destination for the chosen `log_destination_type`.
	   - For an Amazon S3 bucket, specify the key `bucketName` with the name of the bucket and optionally specify the key `prefix` with a path.
	   - For a CloudWatch log group, specify the key `logGroup` with the name of the CloudWatch log group.
	   - For a Kinesis Data Firehose delivery stream, specify the key `deliveryStream` with the name of the delivery stream.
	*/
	LogDestination map[string]string `json:"logDestination,omitempty" yaml:"logDestination,omitempty"`

	// The location to send logs to. Valid values: `S3`, `CloudWatchLogs`, `KinesisDataFirehose`.
	LogDestinationType string `json:"logDestinationType,omitempty" yaml:"logDestinationType,omitempty"`
}
