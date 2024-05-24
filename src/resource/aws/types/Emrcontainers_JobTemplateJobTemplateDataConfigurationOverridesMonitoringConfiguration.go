package types

type Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration struct {
	// Monitoring configurations for CloudWatch.
	CloudWatchMonitoringConfiguration Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationCloudWatchMonitoringConfiguration `json:"cloudWatchMonitoringConfiguration,omitempty" yaml:"cloudWatchMonitoringConfiguration,omitempty"`

	// Monitoring configurations for the persistent application UI.
	PersistentAppUi string `json:"persistentAppUi,omitempty" yaml:"persistentAppUi,omitempty"`

	// Amazon S3 configuration for monitoring log publishing.
	S3MonitoringConfiguration Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfigurationS3MonitoringConfiguration `json:"s3MonitoringConfiguration,omitempty" yaml:"s3MonitoringConfiguration,omitempty"`
}
