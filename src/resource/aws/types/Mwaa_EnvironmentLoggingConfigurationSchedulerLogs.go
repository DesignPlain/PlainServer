package types

type Mwaa_EnvironmentLoggingConfigurationSchedulerLogs struct {
	//
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn,omitempty" yaml:"cloudWatchLogGroupArn,omitempty"`

	// Enabling or disabling the collection of logs
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Logging level. Valid values: `CRITICAL`, `ERROR`, `WARNING`, `INFO`, `DEBUG`. Will be `INFO` by default.
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
}
