package types

type Kinesis_FirehoseDeliveryStreamExtendedS3Configuration struct {
	// The file extension to override the default file extension (for example, `.json`).
	FileExtension string `json:"fileExtension,omitempty" yaml:"fileExtension,omitempty"`

	//
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the S3 bucket
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	// The configuration for dynamic partitioning. Required when using [dynamic partitioning](https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning.html). See `dynamic_partitioning_configuration` block below for details.
	DynamicPartitioningConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDynamicPartitioningConfiguration `json:"dynamicPartitioningConfiguration,omitempty" yaml:"dynamicPartitioningConfiguration,omitempty"`

	// The time zone you prefer. Valid values are `UTC` or a non-3-letter IANA time zones (for example, `America/Los_Angeles`). Default value is `UTC`.
	CustomTimeZone string `json:"customTimeZone,omitempty" yaml:"customTimeZone,omitempty"`

	// Nested argument for the serializer, deserializer, and schema for converting data from the JSON format to the Parquet or ORC format before writing it to Amazon S3. See `data_format_conversion_configuration` block below for details.
	DataFormatConversionConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfiguration `json:"dataFormatConversionConfiguration,omitempty" yaml:"dataFormatConversionConfiguration,omitempty"`

	// The Amazon S3 backup mode.  Valid values are `Disabled` and `Enabled`.  Default value is `Disabled`.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	//
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	// The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, & `HADOOP_SNAPPY`.
	CompressionFormat string `json:"compressionFormat,omitempty" yaml:"compressionFormat,omitempty"`

	// The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The data processing configuration.  See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	//
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
	ErrorOutputPrefix string `json:"errorOutputPrefix,omitempty" yaml:"errorOutputPrefix,omitempty"`

	// The configuration for backup in Amazon S3. Required if `s3_backup_mode` is `Enabled`. Supports the same fields as `s3_configuration` object.
	S3BackupConfiguration Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfiguration `json:"s3BackupConfiguration,omitempty" yaml:"s3BackupConfiguration,omitempty"`

	//
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	/*
	   Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
	   be used.
	*/
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
