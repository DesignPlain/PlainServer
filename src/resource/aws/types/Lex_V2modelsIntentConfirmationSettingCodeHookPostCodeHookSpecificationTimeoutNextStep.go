package types

type Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStep struct {
	// Configuration block for action that the bot executes at runtime when the conversation reaches this step. See `dialog_action`.
	DialogAction Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepDialogAction `json:"dialogAction,omitempty" yaml:"dialogAction,omitempty"`

	// Configuration block for override settings to configure the intent state. See `intent`.
	Intent Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepIntent `json:"intent,omitempty" yaml:"intent,omitempty"`

	// Map of key/value pairs representing session-specific context information. It contains application information passed between Amazon Lex and a client application.
	SessionAttributes map[string]string `json:"sessionAttributes,omitempty" yaml:"sessionAttributes,omitempty"`
}
