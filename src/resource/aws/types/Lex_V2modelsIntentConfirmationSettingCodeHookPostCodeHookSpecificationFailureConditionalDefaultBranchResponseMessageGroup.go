package types

type Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalDefaultBranchResponseMessageGroup struct {
	// Configuration blocks for message variations to send to the user. When variations are defined, Amazon Lex chooses the primary message or one of the variations to send to the user. See `variation`.
	Variations []Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalDefaultBranchResponseMessageGroupVariation `json:"variations,omitempty" yaml:"variations,omitempty"`

	// Configuration block for the primary message that Amazon Lex should send to the user. See `message`.
	Message Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalDefaultBranchResponseMessageGroupMessage `json:"message,omitempty" yaml:"message,omitempty"`
}
