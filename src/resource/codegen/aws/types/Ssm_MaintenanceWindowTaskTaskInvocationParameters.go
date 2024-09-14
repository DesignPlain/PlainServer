package types

type Ssm_MaintenanceWindowTaskTaskInvocationParameters struct {
	// The parameters for an AUTOMATION task type. Documented below.
	AutomationParameters Ssm_MaintenanceWindowTaskTaskInvocationParametersAutomationParameters `json:"automationParameters,omitempty" yaml:"automationParameters,omitempty"`

	// The parameters for a LAMBDA task type. Documented below.
	LambdaParameters Ssm_MaintenanceWindowTaskTaskInvocationParametersLambdaParameters `json:"lambdaParameters,omitempty" yaml:"lambdaParameters,omitempty"`

	// The parameters for a RUN_COMMAND task type. Documented below.
	RunCommandParameters Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters `json:"runCommandParameters,omitempty" yaml:"runCommandParameters,omitempty"`

	// The parameters for a STEP_FUNCTIONS task type. Documented below.
	StepFunctionsParameters Ssm_MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters `json:"stepFunctionsParameters,omitempty" yaml:"stepFunctionsParameters,omitempty"`
}
