package types

type Kinesis_FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration struct {
	//
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// Prefix added to failed records before writing them to S3. Not currently supported for `redshift` destination. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html).
	ErrorOutputPrefix string `json:"errorOutputPrefix,omitempty" yaml:"errorOutputPrefix,omitempty"`

	// The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered S3 files. You can specify an extra prefix to be added in front of the time format prefix. Note that if the prefix ends with a slash, it appears as a folder in the S3 bucket
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The ARN of the S3 bucket
	BucketArn string `json:"bucketArn,omitempty" yaml:"bucketArn,omitempty"`

	//
	BufferingInterval int `json:"bufferingInterval,omitempty" yaml:"bufferingInterval,omitempty"`

	//
	BufferingSize int `json:"bufferingSize,omitempty" yaml:"bufferingSize,omitempty"`

	// The compression format. If no value is specified, the default is `UNCOMPRESSED`. Other supported values are `GZIP`, `ZIP`, `Snappy`, & `HADOOP_SNAPPY`.
	CompressionFormat string `json:"compressionFormat,omitempty" yaml:"compressionFormat,omitempty"`

	/*
	   Specifies the KMS key ARN the stream will use to encrypt data. If not set, no encryption will
	   be used.
	*/
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
