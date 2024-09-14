package types

type Mwaa_EnvironmentLoggingConfiguration struct {
	// Log configuration options for the schedulers. See Module logging configuration for more information. Disabled by default.
	SchedulerLogs Mwaa_EnvironmentLoggingConfigurationSchedulerLogs `json:"schedulerLogs,omitempty" yaml:"schedulerLogs,omitempty"`

	// Log configuration options for DAG tasks. See Module logging configuration for more information. Enabled by default with `INFO` log level.
	TaskLogs Mwaa_EnvironmentLoggingConfigurationTaskLogs `json:"taskLogs,omitempty" yaml:"taskLogs,omitempty"`

	// Log configuration options for the webservers. See Module logging configuration for more information. Disabled by default.
	WebserverLogs Mwaa_EnvironmentLoggingConfigurationWebserverLogs `json:"webserverLogs,omitempty" yaml:"webserverLogs,omitempty"`

	// Log configuration options for the workers. See Module logging configuration for more information. Disabled by default.
	WorkerLogs Mwaa_EnvironmentLoggingConfigurationWorkerLogs `json:"workerLogs,omitempty" yaml:"workerLogs,omitempty"`

	// (Optional) Log configuration options for processing DAGs. See Module logging configuration for more information. Disabled by default.
	DagProcessingLogs Mwaa_EnvironmentLoggingConfigurationDagProcessingLogs `json:"dagProcessingLogs,omitempty" yaml:"dagProcessingLogs,omitempty"`
}
