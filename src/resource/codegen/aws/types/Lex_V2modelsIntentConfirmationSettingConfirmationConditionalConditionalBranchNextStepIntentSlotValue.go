package types

type Lex_V2modelsIntentConfirmationSettingConfirmationConditionalConditionalBranchNextStepIntentSlotValue struct {
	// Value that Amazon Lex determines for the slot. The actual value depends on the setting of the value selection strategy for the bot. You can choose to use the value entered by the user, or you can have Amazon Lex choose the first value in the resolvedValues list.
	InterpretedValue string `json:"interpretedValue,omitempty" yaml:"interpretedValue,omitempty"`
}
