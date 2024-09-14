package lex

import types "libds/aws/types"

type V2modelsSlot struct {
	// Version of the bot associated with the slot.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// Description of the slot.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Determines how slot values are used in Amazon CloudWatch logs. See the `obfuscation_setting` argument reference below.
	ObfuscationSettings []types.Lex_V2modelsSlotObfuscationSetting `json:"obfuscationSettings,omitempty" yaml:"obfuscationSettings,omitempty"`

	// Unique identifier for the slot type associated with this slot.
	SlotTypeId string `json:"slotTypeId,omitempty" yaml:"slotTypeId,omitempty"`

	// Specifications for the constituent sub slots and the expression for the composite slot.
	SubSlotSettings []types.Lex_V2modelsSlotSubSlotSetting `json:"subSlotSettings,omitempty" yaml:"subSlotSettings,omitempty"`

	//
	Timeouts types.Lex_V2modelsSlotTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Identifier of the bot associated with the slot.
	BotId string `json:"botId,omitempty" yaml:"botId,omitempty"`

	// Identifier of the intent that contains the slot.
	IntentId string `json:"intentId,omitempty" yaml:"intentId,omitempty"`

	// Identifier of the language and locale that the slot will be used in.
	LocaleId string `json:"localeId,omitempty" yaml:"localeId,omitempty"`

	// Whether the slot returns multiple values in one response. See the `multiple_values_setting` argument reference below.
	MultipleValuesSettings []types.Lex_V2modelsSlotMultipleValuesSetting `json:"multipleValuesSettings,omitempty" yaml:"multipleValuesSettings,omitempty"`

	// Name of the slot.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Prompts that Amazon Lex sends to the user to elicit a response that provides the value for the slot.

	   The following arguments are optional:
	*/
	ValueElicitationSetting types.Lex_V2modelsSlotValueElicitationSetting `json:"valueElicitationSetting,omitempty" yaml:"valueElicitationSetting,omitempty"`
}
