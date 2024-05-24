package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersAutomationParameters struct {
	// The version of an Automation document to use during task execution.
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameters []Ssm_MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
