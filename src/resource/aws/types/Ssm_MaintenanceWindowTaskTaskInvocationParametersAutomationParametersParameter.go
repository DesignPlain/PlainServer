package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter struct {
	// The parameter name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The array of strings.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
