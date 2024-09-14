package types

type Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationCloudWatchMonitoringConfiguration struct {
	// The name of the log group for log publishing.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The specified name prefix for log streams.
	LogStreamNamePrefix string `json:"logStreamNamePrefix,omitempty" yaml:"logStreamNamePrefix,omitempty"`
}
