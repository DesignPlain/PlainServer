package types

type Kinesis_FirehoseDeliveryStreamHttpEndpointConfigurationCloudwatchLoggingOptions struct {
	// Enables or disables the logging. Defaults to `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The CloudWatch group name for logging. This value is required if `enabled` is true.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The CloudWatch log stream name for logging. This value is required if `enabled` is true.
	LogStreamName string `json:"logStreamName,omitempty" yaml:"logStreamName,omitempty"`
}
