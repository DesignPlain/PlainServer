package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters struct {
	// The inputs for the STEP_FUNCTION task.
	Input string `json:"input,omitempty" yaml:"input,omitempty"`

	// The name of the STEP_FUNCTION task.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
