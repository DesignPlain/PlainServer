package types

type Lex_V2modelsIntentConfirmationSettingDeclinationConditionalDefaultBranch struct {
	// Configuration block for a list of message groups that Amazon Lex uses to respond to the user input. See `response`.
	Response Lex_V2modelsIntentConfirmationSettingDeclinationConditionalDefaultBranchResponse `json:"response,omitempty" yaml:"response,omitempty"`

	// Configuration block for the next step in the conversation. See `next_step`.
	NextStep Lex_V2modelsIntentConfirmationSettingDeclinationConditionalDefaultBranchNextStep `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`
}
