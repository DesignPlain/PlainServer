package types

type Kinesis_FirehoseDeliveryStreamSnowflakeConfiguration struct {
	// The CloudWatch Logging Options for the delivery stream. See `cloudwatch_logging_options` block below for details.
	CloudwatchLoggingOptions Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// The name of the content column.
	ContentColumnName string `json:"contentColumnName,omitempty" yaml:"contentColumnName,omitempty"`

	// The passphrase for the private key.
	KeyPassphrase string `json:"keyPassphrase,omitempty" yaml:"keyPassphrase,omitempty"`

	// The processing configuration. See `processing_configuration` block below for details.
	ProcessingConfiguration Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationProcessingConfiguration `json:"processingConfiguration,omitempty" yaml:"processingConfiguration,omitempty"`

	// The ARN of the IAM role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The S3 backup mode.
	S3BackupMode string `json:"s3BackupMode,omitempty" yaml:"s3BackupMode,omitempty"`

	// The configuration for Snowflake role.
	SnowflakeRoleConfiguration Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSnowflakeRoleConfiguration `json:"snowflakeRoleConfiguration,omitempty" yaml:"snowflakeRoleConfiguration,omitempty"`

	// The URL of the Snowflake account. Format: https://[account_identifier].snowflakecomputing.com.
	AccountUrl string `json:"accountUrl,omitempty" yaml:"accountUrl,omitempty"`

	// The name of the metadata column.
	MetadataColumnName string `json:"metadataColumnName,omitempty" yaml:"metadataColumnName,omitempty"`

	// The private key for authentication. This value is required if `secrets_manager_configuration` is not provided.
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	// The S3 configuration. See `s3_configuration` block below for details.
	S3Configuration Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`

	// The Secrets Manager configuration. See `secrets_manager_configuration` block below for details. This value is required if `user` and `private_key` are not provided.
	SecretsManagerConfiguration Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSecretsManagerConfiguration `json:"secretsManagerConfiguration,omitempty" yaml:"secretsManagerConfiguration,omitempty"`

	// The VPC configuration for Snowflake.
	SnowflakeVpcConfiguration Kinesis_FirehoseDeliveryStreamSnowflakeConfigurationSnowflakeVpcConfiguration `json:"snowflakeVpcConfiguration,omitempty" yaml:"snowflakeVpcConfiguration,omitempty"`

	// The Snowflake database name.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The Snowflake table name.
	Table string `json:"table,omitempty" yaml:"table,omitempty"`

	// After an initial failure to deliver to Snowflake, the total amount of time, in seconds between 0 to 7200, during which Firehose re-attempts delivery (including the first attempt).  After this time has elapsed, the failed documents are written to Amazon S3.  The default value is 60s.  There will be no retry if the value is 0.
	RetryDuration int `json:"retryDuration,omitempty" yaml:"retryDuration,omitempty"`

	// The Snowflake schema name.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	// The user for authentication. This value is required if `secrets_manager_configuration` is not provided.
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// The data loading option.
	DataLoadingOption string `json:"dataLoadingOption,omitempty" yaml:"dataLoadingOption,omitempty"`
}
