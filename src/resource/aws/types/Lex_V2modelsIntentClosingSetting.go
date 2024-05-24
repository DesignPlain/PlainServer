package types

type Lex_V2modelsIntentClosingSetting struct {
	// Whether an intent's closing response is used. When this field is false, the closing response isn't sent to the user. If the active field isn't specified, the default is true.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Configuration block for response that Amazon Lex sends to the user when the intent is complete. See `closing_response`.
	ClosingResponse Lex_V2modelsIntentClosingSettingClosingResponse `json:"closingResponse,omitempty" yaml:"closingResponse,omitempty"`

	// Configuration block for list of conditional branches associated with the intent's closing response. These branches are executed when the `next_step` attribute is set to `EvalutateConditional`. See `conditional`.
	Conditional Lex_V2modelsIntentClosingSettingConditional `json:"conditional,omitempty" yaml:"conditional,omitempty"`

	// Next step that the bot executes after playing the intent's closing response. See `next_step`.
	NextStep Lex_V2modelsIntentClosingSettingNextStep `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`
}
