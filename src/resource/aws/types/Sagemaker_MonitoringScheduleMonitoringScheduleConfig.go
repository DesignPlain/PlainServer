package types

type Sagemaker_MonitoringScheduleMonitoringScheduleConfig struct {
	// The name of the monitoring job definition to schedule.
	MonitoringJobDefinitionName string `json:"monitoringJobDefinitionName,omitempty" yaml:"monitoringJobDefinitionName,omitempty"`

	// The type of the monitoring job definition to schedule. Valid values are `DataQuality`, `ModelQuality`, `ModelBias` or `ModelExplainability`
	MonitoringType string `json:"monitoringType,omitempty" yaml:"monitoringType,omitempty"`

	// Configures the monitoring schedule. Fields are documented below.
	ScheduleConfig Sagemaker_MonitoringScheduleMonitoringScheduleConfigScheduleConfig `json:"scheduleConfig,omitempty" yaml:"scheduleConfig,omitempty"`
}
