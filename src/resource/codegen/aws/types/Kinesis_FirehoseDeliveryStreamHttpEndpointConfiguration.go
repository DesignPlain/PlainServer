package types

type Kinesis_FirehoseDeliveryStreamHttpEndpointConfiguration struct {
	// Total amount of seconds Firehose spends on retries. This duration starts after the initial attempt fails, It does not include the time periods during which Firehose waits for acknowledgment from the specified destination after each attempt. Valid values between `0` and `7200`. Default is `300`.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// Defines how documents should be delivered to Amazon S3.  Valid values are `FailedDataOnly` and `AllData`.  Default value is `FailedDataOnly`.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	// The S3 Configuration. See `s3_configuration` block below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// The HTTP endpoint URL to which Kinesis Firehose sends your data.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// The access key required for Kinesis Firehose to authenticate with the HTTP endpoint selected as the destination.
	AccessKey string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`

	// Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination. The default value is 300 (5 minutes).
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	// The HTTP endpoint name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The data processing configuration.  See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The Secret Manager Configuration. See `secrets_manager_configuration` block below for details.
	SecretsManagerConfiguration Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationSecretsManagerConfiguration `json:"secretsManagerConfiguration,omitempty" yaml:"secretsManagerConfiguration,omitempty"`

	// Buffer incoming data to the specified size, in MBs, before delivering it to the destination. The default value is 5.
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// The request configuration.  See `request_configuration` block below for details.
	RequestConfiguration Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationRequestConfiguration `json:"requestConfiguration,omitempty" yaml:"requestConfiguration,omitempty"`

	// Kinesis Data Firehose uses this IAM role for all the permissions that the delivery stream needs. The pattern needs to be `arn:.-`.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
