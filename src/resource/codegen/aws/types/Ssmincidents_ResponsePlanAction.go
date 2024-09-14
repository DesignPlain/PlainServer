package types

type Ssmincidents_ResponsePlanAction struct {
	// The Systems Manager automation document to start as the runbook at the beginning of the incident. The following values are supported:
	SsmAutomations []Ssmincidents_ResponsePlanActionSsmAutomation `json:"ssmAutomations,omitempty" yaml:"ssmAutomations,omitempty"`
}
