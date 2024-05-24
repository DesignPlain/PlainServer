package types

type Kinesis_FirehoseDeliveryStreamRedshiftConfiguration struct {
	// The configuration for backup in Amazon S3. Required if `s3_backup_mode` is `Enabled`. Supports the same fields as `s3_configuration` object.
	S3BackupConfiguration Kinesis_FirehoseDeliveryStreamRedshiftConfigurationS3BackupConfiguration `json:"s3BackupConfiguration,omitempty" yaml:"s3BackupConfiguration,omitempty"`

	// The jdbcurl of the redshift cluster.
	ClusterJdbcurl string `json:"clusterJdbcurl,omitempty" yaml:"clusterJdbcurl,omitempty"`

	// The name of the table in the redshift cluster that the s3 bucket will copy to.
	DataTableName string `json:"dataTableName,omitempty" yaml:"dataTableName,omitempty"`

	// The password for the username above.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The data processing configuration.  See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamRedshiftConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The length of time during which Firehose retries delivery after a failure, starting from the initial request and including the first attempt. The default value is 3600 seconds (60 minutes). Firehose does not retry if the value of DurationInSeconds is 0 (zero) or if the first delivery attempt takes longer than the current value.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// The arn of the role the stream assumes.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The Amazon S3 backup mode.  Valid values are `Disabled` and `Enabled`.  Default value is `Disabled`.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	// The S3 Configuration. See s3_configuration below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamRedshiftConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamRedshiftConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// Copy options for copying the data from the s3 intermediate bucket into redshift, for example to change the default delimiter. For valid values, see the [AWS documentation](http://docs.aws.amazon.com/firehose/latest/APIReference/API_CopyCommand.html)
	CopyOptions string `json:"copyOptions,omitempty" yaml:"copyOptions,omitempty"`

	// The data table columns that will be targeted by the copy command.
	DataTableColumns string `json:"dataTableColumns,omitempty" yaml:"dataTableColumns,omitempty"`

	// The username that the firehose delivery stream will assume. It is strongly recommended that the username and password provided is used exclusively for Amazon Kinesis Firehose purposes, and that the permissions for the account are restricted for Amazon Redshift INSERT permissions.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
