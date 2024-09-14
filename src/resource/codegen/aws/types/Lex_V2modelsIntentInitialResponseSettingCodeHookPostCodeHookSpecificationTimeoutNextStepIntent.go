package types

type Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStepIntent struct {
	// Name of the intent.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for all of the slot value overrides for the intent. The name of the slot maps to the value of the slot. Slots that are not included in the map aren't overridden. See `slot`.
	Slots []Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStepIntentSlot `json:"slots,omitempty" yaml:"slots,omitempty"`
}
