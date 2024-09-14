package kinesisanalyticsv2

import types "libds/aws/types"

type Application struct {
	// The application's configuration
	ApplicationConfiguration types.Kinesisanalyticsv2_ApplicationApplicationConfiguration `json:"applicationConfiguration,omitempty" yaml:"applicationConfiguration,omitempty"`

	// A CloudWatch log stream to monitor application configuration errors.
	CloudwatchLoggingOptions types.Kinesisanalyticsv2_ApplicationCloudwatchLoggingOptions `json:"cloudwatchLoggingOptions,omitempty" yaml:"cloudwatchLoggingOptions,omitempty"`

	// A summary description of the application.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether to force stop an unresponsive Flink-based application.
	ForceStop bool `json:"forceStop,omitempty" yaml:"forceStop,omitempty"`

	// Whether to start or stop the application.
	StartApplication bool `json:"startApplication,omitempty" yaml:"startApplication,omitempty"`

	// A map of tags to assign to the application. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The application's mode. Valid values are `STREAMING`, `INTERACTIVE`.
	ApplicationMode string `json:"applicationMode,omitempty" yaml:"applicationMode,omitempty"`

	// The name of the application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`, `FLINK-1_11`, `FLINK-1_13`, `FLINK-1_15`, `FLINK-1_18`, `FLINK-1_19`.
	RuntimeEnvironment string `json:"runtimeEnvironment,omitempty" yaml:"runtimeEnvironment,omitempty"`

	// The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole string `json:"serviceExecutionRole,omitempty" yaml:"serviceExecutionRole,omitempty"`
}
