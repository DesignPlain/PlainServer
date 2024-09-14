package types

type Securityhub_AutomationRuleActionFindingFieldsUpdateNote struct {
	// The principal that updated the note.
	UpdatedBy string `json:"updatedBy,omitempty" yaml:"updatedBy,omitempty"`

	// The updated note text.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}
