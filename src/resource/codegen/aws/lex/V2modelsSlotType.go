package lex

import types "libds/aws/types"

type V2modelsSlotType struct {
	// Identifier of the language and locale where this slot type is used. All of the bots, slot types, and slots used by the intent must have the same locale.
	LocaleId string `json:"localeId,omitempty" yaml:"localeId,omitempty"`

	// List of SlotTypeValue objects that defines the values that the slot type can take. Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for a slot. See `slot_type_values` argument reference below.
	SlotTypeValues types.Lex_V2modelsSlotTypeSlotTypeValues `json:"slotTypeValues,omitempty" yaml:"slotTypeValues,omitempty"`

	// Determines the strategy that Amazon Lex uses to select a value from the list of possible values. The field can be set to one of the following values: `ORIGINAL_VALUE` returns the value entered by the user, if the user value is similar to the slot value. `TOP_RESOLUTION` if there is a resolution list for the slot, return the first value in the resolution list. If there is no resolution list, return null. If you don't specify the valueSelectionSetting parameter, the default is ORIGINAL_VALUE. See `value_selection_setting` argument reference below.
	ValueSelectionSetting types.Lex_V2modelsSlotTypeValueSelectionSetting `json:"valueSelectionSetting,omitempty" yaml:"valueSelectionSetting,omitempty"`

	// Identifier of the bot associated with this slot type.
	BotId string `json:"botId,omitempty" yaml:"botId,omitempty"`

	// Specifications for a composite slot type. See `composite_slot_type_setting` argument reference below.
	CompositeSlotTypeSetting types.Lex_V2modelsSlotTypeCompositeSlotTypeSetting `json:"compositeSlotTypeSetting,omitempty" yaml:"compositeSlotTypeSetting,omitempty"`

	// Description of the slot type.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Built-in slot type used as a parent of this slot type. When you define a parent slot type, the new slot type has the configuration of the parent slot type. Only AMAZON.AlphaNumeric is supported.
	ParentSlotTypeSignature string `json:"parentSlotTypeSignature,omitempty" yaml:"parentSlotTypeSignature,omitempty"`

	//
	Timeouts types.Lex_V2modelsSlotTypeTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Version of the bot associated with this slot type.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// Type of external information used to create the slot type. See `external_source_setting` argument reference below.
	ExternalSourceSetting types.Lex_V2modelsSlotTypeExternalSourceSetting `json:"externalSourceSetting,omitempty" yaml:"externalSourceSetting,omitempty"`

	/*
	   Name of the slot type

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
