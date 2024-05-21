package types

type Dataproc_WorkflowTemplateJobSparkJobLoggingConfig struct {
	// The per-package log levels for the driver. This may include "root" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" yaml:"driverLogLevels,omitempty"`
}
