package iot

type LoggingOptions struct {
	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The default logging level. Valid Values: `"DEBUG"`, `"INFO"`, `"ERROR"`, `"WARN"`, `"DISABLED"`.
	DefaultLogLevel string `json:"defaultLogLevel,omitempty" yaml:"defaultLogLevel,omitempty"`

	// If `true` all logs are disabled. The default is `false`.
	DisableAllLogs bool `json:"disableAllLogs,omitempty" yaml:"disableAllLogs,omitempty"`
}
