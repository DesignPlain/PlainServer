package types

type Kinesis_FirehoseDeliveryStreamSplunkConfiguration struct {
	// After an initial failure to deliver to Splunk, the total amount of time, in seconds between 0 to 7200, during which Firehose re-attempts delivery (including the first attempt).  After this time has elapsed, the failed documents are written to Amazon S3.  The default value is 300s.  There will be no retry if the value is 0.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// The S3 Configuration. See `s3_configuration` block below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamSplunkConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// Buffer incoming data for the specified period of time, in seconds between 0 to 60, before delivering it to the destination.  The default value is 60s.
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	// Buffer incoming data to the specified size, in MBs between 1 to 5, before delivering it to the destination.  The default value is 5MB.
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamSplunkConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// The HTTP Event Collector (HEC) endpoint to which Kinesis Firehose sends your data.
	HecEndpoint string `json:"hecEndpoint,omitempty" yaml:"hecEndpoint,omitempty"`

	/*
	   Defines how documents should be delivered to Amazon S3.  Valid values are `FailedEventsOnly` and `AllEvents`.  Default value is `FailedEventsOnly`.
	   `secrets_manager_configuration` - (Optional) The Secrets Manager configuration. See `secrets_manager_configuration` block below for details. This value is required if `hec_token` is not provided.
	*/
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	//
	SecretsManagerConfiguration Kinesis_FirehoseDeliveryStreamSplunkConfigurationSecretsManagerConfiguration `json:"secretsManagerConfiguration,omitempty" yaml:"secretsManagerConfiguration,omitempty"`

	// The amount of time, in seconds between 180 and 600, that Kinesis Firehose waits to receive an acknowledgment from Splunk after it sends it data.
	HecAcknowledgmentTimeout int `json:"hecAcknowledgmentTimeout,omitempty" yaml:"hecAcknowledgmentTimeout,omitempty"`

	// The HEC endpoint type. Valid values are `Raw` or `Event`. The default value is `Raw`.
	HecEndpointType string `json:"hecEndpointType,omitempty" yaml:"hecEndpointType,omitempty"`

	// The GUID that you obtain from your Splunk cluster when you create a new HEC endpoint. This value is required if `secrets_manager_configuration` is not provided.
	HecToken string `json:"hecToken,omitempty" yaml:"hecToken,omitempty"`

	// The data processing configuration.  See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`
}
