package types



type Composer_EnvironmentConfigDataRetentionConfig struct {
	// Optional. The configuration setting for Task Logs.
	TaskLogsRetentionConfigs []Composer_EnvironmentConfigDataRetentionConfigTaskLogsRetentionConfig `json:"taskLogsRetentionConfigs,omitempty" yaml:"taskLogsRetentionConfigs,omitempty"`
}
