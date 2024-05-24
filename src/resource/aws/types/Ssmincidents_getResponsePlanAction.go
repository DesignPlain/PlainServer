package types

type Ssmincidents_getResponsePlanAction struct {
	// The Systems Manager automation document to start as the runbook at the beginning of the incident. The following values are supported:
	SsmAutomations []Ssmincidents_getResponsePlanActionSsmAutomation `json:"ssmAutomations,omitempty" yaml:"ssmAutomations,omitempty"`
}
