package types

type Verifiedaccess_InstanceLoggingConfigurationAccessLogsCloudwatchLogs struct {
	// Indicates whether logging is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the CloudWatch Logs Log Group.
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
