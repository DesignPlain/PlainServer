package types

type Securityhub_AutomationRuleActionFindingFieldsUpdateSeverity struct {
	// The severity value of the finding. The allowed values are the following `INFORMATIONAL`, `LOW`, `MEDIUM`, `HIGH` and `CRITICAL`.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`

	// The native severity as defined by the AWS service or integrated partner product that generated the finding.
	Product float64 `json:"product,omitempty" yaml:"product,omitempty"`
}
