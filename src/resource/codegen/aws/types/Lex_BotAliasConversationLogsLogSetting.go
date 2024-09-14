package types

type Lex_BotAliasConversationLogsLogSetting struct {
	// The destination where logs are delivered. Options are `CLOUDWATCH_LOGS` or `S3`.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The Amazon Resource Name (ARN) of the key used to encrypt audio logs in an S3 bucket. This can only be specified when `destination` is set to `S3`. Must be between 20 and 2048 characters in length.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The type of logging that is enabled. Options are `AUDIO` or `TEXT`.
	LogType string `json:"logType,omitempty" yaml:"logType,omitempty"`

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or S3 bucket where the logs are delivered. Must be less than or equal to 2048 characters in length.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The prefix of the S3 object key for `AUDIO` logs or the log stream name for `TEXT` logs.
	ResourcePrefix string `json:"resourcePrefix,omitempty" yaml:"resourcePrefix,omitempty"`
}
