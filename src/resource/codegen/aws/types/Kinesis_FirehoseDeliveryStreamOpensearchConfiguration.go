package types

type Kinesis_FirehoseDeliveryStreamOpensearchConfiguration struct {
	// The ARN of the Amazon ES domain.  The pattern needs to be `arn:.-`.  Conflicts with `cluster_endpoint`.
	DomainArn string `json:"domainArn,omitempty" yaml:"domainArn,omitempty"`

	// The OpenSearch index rotation period.  Index rotation appends a timestamp to the IndexName to facilitate expiration of old data.  Valid values are `NoRotation`, `OneHour`, `OneDay`, `OneWeek`, and `OneMonth`.  The default value is `OneDay`.
	IndexRotationPeriod string `json:"indexRotationPeriod,omitempty" yaml:"indexRotationPeriod,omitempty"`

	// Buffer incoming data for the specified period of time, in seconds between 0 to 900, before delivering it to the destination.  The default value is 300s.
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	// The endpoint to use when communicating with the cluster. Conflicts with `domain_arn`.
	ClusterEndpoint string `json:"clusterEndpoint,omitempty" yaml:"clusterEndpoint,omitempty"`

	// The data processing configuration. See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamOpensearchConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The S3 Configuration. See `s3_configuration` block below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamOpensearchConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// The Elasticsearch type name with maximum length of 100 characters. Types are deprecated in OpenSearch_1.1. TypeName must be empty.
	TypeName string `json:"typeName,omitempty" yaml:"typeName,omitempty"`

	// Buffer incoming data to the specified size, in MBs between 1 to 100, before delivering it to the destination.  The default value is 5MB.
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// The OpenSearch index name.
	IndexName string `json:"indexName,omitempty" yaml:"indexName,omitempty"`

	// The ARN of the IAM role to be assumed by Firehose for calling the Amazon ES Configuration API and for indexing documents.  The IAM role must have permission for `DescribeDomain`, `DescribeDomains`, and `DescribeDomainConfig`.  The pattern needs to be `arn:.-`.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamOpensearchConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// The method for setting up document ID. See [`document_id_options` block] below for details.
	DocumentIdOptions Kinesis_FirehoseDeliveryStreamOpensearchConfigurationDocumentIdOptions `json:"documentIdOptions,omitempty" yaml:"documentIdOptions,omitempty"`

	// The VPC configuration for the delivery stream to connect to OpenSearch associated with the VPC. See `vpc_config` block below for details.
	VpcConfig Kinesis_FirehoseDeliveryStreamOpensearchConfigurationVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// After an initial failure to deliver to Amazon OpenSearch, the total amount of time, in seconds between 0 to 7200, during which Firehose re-attempts delivery (including the first attempt).  After this time has elapsed, the failed documents are written to Amazon S3.  The default value is 300s.  There will be no retry if the value is 0.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// Defines how documents should be delivered to Amazon S3.  Valid values are `FailedDocumentsOnly` and `AllDocuments`.  Default value is `FailedDocumentsOnly`.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`
}
