package types

type Lex_V2modelsSlotTypeSlotTypeValues struct {
	// Value of the slot type entry.  See `sample_value` argument reference below.
	SampleValues []Lex_V2modelsSlotTypeSlotTypeValuesSampleValue `json:"sampleValues,omitempty" yaml:"sampleValues,omitempty"`

	// Additional values related to the slot type entry. See `sample_value` argument reference below.
	Synonyms []Lex_V2modelsSlotTypeSlotTypeValuesSynonym `json:"synonyms,omitempty" yaml:"synonyms,omitempty"`
}
