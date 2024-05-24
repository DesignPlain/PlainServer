package types

type Securityhub_AutomationRuleActionFindingFieldsUpdate struct {
	// The rule action updates the `Criticality` field of a finding.
	Criticality int `json:"criticality,omitempty" yaml:"criticality,omitempty"`

	// The rule action updates the `VerificationState` field of a finding. The allowed values are the following `UNKNOWN`, `TRUE_POSITIVE`, `FALSE_POSITIVE` and `BENIGN_POSITIVE`.
	VerificationState string `json:"verificationState,omitempty" yaml:"verificationState,omitempty"`

	// A resource block that is used to update information about the investigation into the finding. Documented below.
	Workflow Securityhub_AutomationRuleActionFindingFieldsUpdateWorkflow `json:"workflow,omitempty" yaml:"workflow,omitempty"`

	// The rule action updates the `Confidence` field of a finding.
	Confidence int `json:"confidence,omitempty" yaml:"confidence,omitempty"`

	// A resource block that updates the note. Documented below.
	Note Securityhub_AutomationRuleActionFindingFieldsUpdateNote `json:"note,omitempty" yaml:"note,omitempty"`

	// A resource block that the rule action updates the `RelatedFindings` field of a finding. Documented below.
	RelatedFindings []Securityhub_AutomationRuleActionFindingFieldsUpdateRelatedFinding `json:"relatedFindings,omitempty" yaml:"relatedFindings,omitempty"`

	// A resource block that updates to the severity information for a finding. Documented below.
	Severity Securityhub_AutomationRuleActionFindingFieldsUpdateSeverity `json:"severity,omitempty" yaml:"severity,omitempty"`

	// The rule action updates the `Types` field of a finding.
	Types []string `json:"types,omitempty" yaml:"types,omitempty"`

	// The rule action updates the `UserDefinedFields` field of a finding.
	UserDefinedFields map[string]string `json:"userDefinedFields,omitempty" yaml:"userDefinedFields,omitempty"`
}
