package types

type Opsworks_CustomLayerCloudwatchConfigurationLogStream struct {
	// Specifies how the timestamp is extracted from logs. For more information, see the CloudWatch Logs Agent Reference (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AgentReference.html).
	DatetimeFormat string `json:"datetimeFormat,omitempty" yaml:"datetimeFormat,omitempty"`

	// Specifies log files that you want to push to CloudWatch Logs. File can point to a specific file or multiple files (by using wild card characters such as /var/log/system.log-).
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	// Specifies where to start to read data (`start_of_file` or `end_of_file`). The default is `start_of_file`.
	InitialPosition string `json:"initialPosition,omitempty" yaml:"initialPosition,omitempty"`

	// Specifies the pattern for identifying the start of a log message.
	MultilineStartPattern string `json:"multilineStartPattern,omitempty" yaml:"multilineStartPattern,omitempty"`

	// Specifies the time zone of log event time stamps.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// Specifies the max number of log events in a batch, up to `10000`. The default value is `1000`.
	BatchCount int `json:"batchCount,omitempty" yaml:"batchCount,omitempty"`

	// Specifies the maximum size of log events in a batch, in bytes, up to `1048576` bytes. The default value is `32768` bytes.
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	// Specifies the time duration for the batching of log events. The minimum value is `5000` and default value is `5000`.
	BufferDuration int `json:"bufferDuration,omitempty" yaml:"bufferDuration,omitempty"`

	// Specifies the encoding of the log file so that the file can be read correctly. The default is `utf_8`.
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	// Specifies the range of lines for identifying a file. The valid values are one number, or two dash-delimited numbers, such as `1`, `2-5`. The default value is `1`.
	FileFingerprintLines string `json:"fileFingerprintLines,omitempty" yaml:"fileFingerprintLines,omitempty"`

	// Specifies the destination log group. A log group is created automatically if it doesn't already exist.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`
}
