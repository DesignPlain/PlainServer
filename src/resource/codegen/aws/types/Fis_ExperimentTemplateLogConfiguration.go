package types

type Fis_ExperimentTemplateLogConfiguration struct {
	// The configuration for experiment logging to Amazon CloudWatch Logs. See below.
	CloudwatchLogsConfiguration Fis_ExperimentTemplateLogConfigurationCloudwatchLogsConfiguration `json:"cloudwatchLogsConfiguration,omitempty" yaml:"cloudwatchLogsConfiguration,omitempty"`

	// The schema version. See [documentation](https://docs.aws.amazon.com/fis/latest/userguide/monitoring-logging.html#experiment-log-schema) for the list of schema versions.
	LogSchemaVersion int `json:"logSchemaVersion,omitempty" yaml:"logSchemaVersion,omitempty"`

	// The configuration for experiment logging to Amazon S3. See below.
	S3Configuration Fis_ExperimentTemplateLogConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`
}
