package types

type Ssmincidents_ResponsePlanActionSsmAutomationParameter struct {
	// The name of the response plan.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The values for the associated parameter name.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
