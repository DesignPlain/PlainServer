package types

type Composer_getEnvironmentConfigDataRetentionConfig struct {
	// Optional. The configuration setting for Task Logs.
	TaskLogsRetentionConfigs []Composer_getEnvironmentConfigDataRetentionConfigTaskLogsRetentionConfig `json:"taskLogsRetentionConfigs,omitempty" yaml:"taskLogsRetentionConfigs,omitempty"`
}
