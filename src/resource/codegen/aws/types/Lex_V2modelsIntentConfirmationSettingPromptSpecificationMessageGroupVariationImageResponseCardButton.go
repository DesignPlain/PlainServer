package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationMessageGroupVariationImageResponseCardButton struct {
	// Value returned to Amazon Lex when the user chooses this button. This must be one of the slot values configured for the slot.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Text that appears on the button. Use this to tell the user what value is returned when they choose this button.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}
