package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfiguration struct {
	// Describes whether to use the default CloudWatch logging configuration for an application. Valid values: `CUSTOM`, `DEFAULT`. Set this attribute to `CUSTOM` in order for any specified `log_level` or `metrics_level` attribute values to be effective.
	ConfigurationType string `json:"configurationType,omitempty" yaml:"configurationType,omitempty"`

	// Describes the verbosity of the CloudWatch Logs for an application. Valid values: `DEBUG`, `ERROR`, `INFO`, `WARN`.
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`

	// Describes the granularity of the CloudWatch Logs for an application. Valid values: `APPLICATION`, `OPERATOR`, `PARALLELISM`, `TASK`.
	MetricsLevel string `json:"metricsLevel,omitempty" yaml:"metricsLevel,omitempty"`
}
