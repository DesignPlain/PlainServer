package types

type Lex_V2modelsIntentInitialResponseSetting struct {
	// Configuration block for conditional branches. Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition. See `conditional`.
	Conditional Lex_V2modelsIntentInitialResponseSettingConditional `json:"conditional,omitempty" yaml:"conditional,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `initial_response`.
	InitialResponse Lex_V2modelsIntentInitialResponseSettingInitialResponse `json:"initialResponse,omitempty" yaml:"initialResponse,omitempty"`

	// Configuration block for the next step in the conversation. See `next_step`.
	NextStep Lex_V2modelsIntentInitialResponseSettingNextStep `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`

	// Configuration block for the dialog code hook that is called by Amazon Lex at a step of the conversation. See `code_hook`.
	CodeHook Lex_V2modelsIntentInitialResponseSettingCodeHook `json:"codeHook,omitempty" yaml:"codeHook,omitempty"`
}
