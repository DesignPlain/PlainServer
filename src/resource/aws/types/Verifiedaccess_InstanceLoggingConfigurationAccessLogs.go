package types

type Verifiedaccess_InstanceLoggingConfigurationAccessLogs struct {
	// A block that specifies configures sending Verified Access logs to CloudWatch Logs. Detailed below.
	CloudwatchLogs Verifiedaccess_InstanceLoggingConfigurationAccessLogsCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// Include trust data sent by trust providers into the logs.
	IncludeTrustContext bool `json:"includeTrustContext,omitempty" yaml:"includeTrustContext,omitempty"`

	// A block that specifies configures sending Verified Access logs to Kinesis. Detailed below.
	KinesisDataFirehose Verifiedaccess_InstanceLoggingConfigurationAccessLogsKinesisDataFirehose `json:"kinesisDataFirehose,omitempty" yaml:"kinesisDataFirehose,omitempty"`

	// The logging version to use. Refer to [VerifiedAccessLogOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessLogOptions.html) for the allowed values.
	LogVersion string `json:"logVersion,omitempty" yaml:"logVersion,omitempty"`

	// A block that specifies configures sending Verified Access logs to S3. Detailed below.
	S3 Verifiedaccess_InstanceLoggingConfigurationAccessLogsS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
