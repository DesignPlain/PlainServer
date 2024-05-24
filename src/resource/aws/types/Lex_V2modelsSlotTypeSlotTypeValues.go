package types

type Lex_V2modelsSlotTypeSlotTypeValues struct {
	// List of SlotTypeValue objects that defines the values that the slot type can take. Each value can have a list of synonyms, additional values that help train the machine learning model about the values that it resolves for a slot. See `slot_type_values` argument reference below.
	SlotTypeValues []string `json:"slotTypeValues,omitempty" yaml:"slotTypeValues,omitempty"`

	// Additional values related to the slot type entry. See `sample_value` argument reference below.
	Synonyms []Lex_V2modelsSlotTypeSlotTypeValuesSynonym `json:"synonyms,omitempty" yaml:"synonyms,omitempty"`
}
