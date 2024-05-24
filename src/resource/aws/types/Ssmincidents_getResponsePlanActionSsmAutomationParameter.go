package types

type Ssmincidents_getResponsePlanActionSsmAutomationParameter struct {
	// The values for the associated parameter name.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The name of the PagerDuty configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
