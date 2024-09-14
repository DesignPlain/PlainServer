package types

type Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchNextStepIntentSlot struct {
	// Which attempt to configure. Valid values are `Initial`, `Retry1`, `Retry2`, `Retry3`, `Retry4`, `Retry5`.
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	// When the shape value is `List`, `values` contains a list of slot values. When the value is `Scalar`, `value` contains a single value.
	Shape string `json:"shape,omitempty" yaml:"shape,omitempty"`

	// Configuration block for the current value of the slot. See `value`.
	Value Lex_V2modelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchNextStepIntentSlotValue `json:"value,omitempty" yaml:"value,omitempty"`
}
