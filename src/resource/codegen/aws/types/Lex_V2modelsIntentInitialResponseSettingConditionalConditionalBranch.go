package types

type Lex_V2modelsIntentInitialResponseSettingConditionalConditionalBranch struct {
	// Configuration block for the next step in the conversation. See `next_step`.
	NextStep Lex_V2modelsIntentInitialResponseSettingConditionalConditionalBranchNextStep `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`

	// Configuration block for a list of message groups that Amazon Lex uses to respond to the user input. See `response`.
	Response Lex_V2modelsIntentInitialResponseSettingConditionalConditionalBranchResponse `json:"response,omitempty" yaml:"response,omitempty"`

	// Configuration block for the expression to evaluate. If the condition is true, the branch's actions are taken. See `condition`.
	Condition Lex_V2modelsIntentInitialResponseSettingConditionalConditionalBranchCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Name of the branch.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
