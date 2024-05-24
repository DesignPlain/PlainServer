package types

type Appsync_GraphQLApiLogConfig struct {
	// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
	CloudwatchLogsRoleArn string `json:"cloudwatchLogsRoleArn,omitempty" yaml:"cloudwatchLogsRoleArn,omitempty"`

	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: `true`, `false`. Default value: `false`
	ExcludeVerboseContent bool `json:"excludeVerboseContent,omitempty" yaml:"excludeVerboseContent,omitempty"`

	// Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
	FieldLogLevel string `json:"fieldLogLevel,omitempty" yaml:"fieldLogLevel,omitempty"`
}
