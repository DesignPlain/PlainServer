package types

type Securityhub_AutomationRuleActionFindingFieldsUpdateWorkflow struct {
	// The status of the investigation into the finding. The allowed values are the following `NEW`, `NOTIFIED`, `RESOLVED` and `SUPPRESSED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
