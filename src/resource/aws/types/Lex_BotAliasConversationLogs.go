package types

type Lex_BotAliasConversationLogs struct {
	// The Amazon Resource Name (ARN) of the IAM role used to write your logs to CloudWatch Logs or an S3 bucket. Must be between 20 and 2048 characters in length.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// The settings for your conversation logs. You can log text, audio, or both. Attributes are documented under log_settings.
	LogSettings []Lex_BotAliasConversationLogsLogSetting `json:"logSettings,omitempty" yaml:"logSettings,omitempty"`
}
