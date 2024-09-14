package types

type Lex_V2modelsSlotSubSlotSettingSlotSpecification struct {
	//
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	// Unique identifier for the slot type associated with this slot.
	SlotTypeId string `json:"slotTypeId,omitempty" yaml:"slotTypeId,omitempty"`

	/*
	   Prompts that Amazon Lex sends to the user to elicit a response that provides the value for the slot.

	   The following arguments are optional:
	*/
	ValueElicitationSettings []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting `json:"valueElicitationSettings,omitempty" yaml:"valueElicitationSettings,omitempty"`
}
