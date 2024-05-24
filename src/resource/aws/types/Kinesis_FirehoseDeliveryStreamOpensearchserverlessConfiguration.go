package types

type Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfiguration struct {
	// After an initial failure to deliver to the Serverless offering for Amazon OpenSearch Service, the total amount of time, in seconds between 0 to 7200, during which Kinesis Data Firehose retries delivery (including the first attempt).  After this time has elapsed, the failed documents are written to Amazon S3.  The default value is 300s.  There will be no retry if the value is 0.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// The S3 Configuration. See `s3_configuration` block below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// The VPC configuration for the delivery stream to connect to OpenSearch Serverless associated with the VPC. See `vpc_config` block below for details.
	VpcConfig Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfigurationVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Buffer incoming data for the specified period of time, in seconds between 0 to 900, before delivering it to the destination.  The default value is 300s.
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	// The data processing configuration.  See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The endpoint to use when communicating with the collection in the Serverless offering for Amazon OpenSearch Service.
	CollectionEndpoint string `json:"collectionEndpoint,omitempty" yaml:"collectionEndpoint,omitempty"`

	// The Serverless offering for Amazon OpenSearch Service index name.
	IndexName string `json:"indexName,omitempty" yaml:"indexName,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose for calling the Serverless offering for Amazon OpenSearch Service Configuration API and for indexing documents.  The pattern needs to be `arn:.-`.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Defines how documents should be delivered to Amazon S3.  Valid values are `FailedDocumentsOnly` and `AllDocuments`.  Default value is `FailedDocumentsOnly`.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	// Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamOpensearchserverlessConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`
}
