package types

type Securityhub_AutomationRuleActionFindingFieldsUpdateNote struct {
	// The updated note text.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`

	// The principal that updated the note.
	UpdatedBy string `json:"updatedBy,omitempty" yaml:"updatedBy,omitempty"`
}
