package types

type Emrcontainers_JobTemplateJobTemplateDataConfigurationOverrides struct {
	// The configurations for the application running by the job run.
	ApplicationConfigurations []Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesApplicationConfiguration `json:"applicationConfigurations,omitempty" yaml:"applicationConfigurations,omitempty"`

	// The configurations for monitoring.
	MonitoringConfiguration Emrcontainers_JobTemplateJobTemplateDataConfigurationOverridesMonitoringConfiguration `json:"monitoringConfiguration,omitempty" yaml:"monitoringConfiguration,omitempty"`
}
