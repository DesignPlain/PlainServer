package types

type Securityhub_AutomationRuleAction struct {
	// A block that specifies that the automation rule action is an update to a finding field.  Documented below.
	FindingFieldsUpdate Securityhub_AutomationRuleActionFindingFieldsUpdate `json:"findingFieldsUpdate,omitempty" yaml:"findingFieldsUpdate,omitempty"`

	// Specifies that the rule action should update the `Types` finding field. The `Types` finding field classifies findings in the format of namespace/category/classifier.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
