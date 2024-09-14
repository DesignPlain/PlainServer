package types

type Lex_SlotTypeEnumerationValue struct {
	// Additional values related to the slot type value. Each item must be less than or equal to 140 characters in length.
	Synonyms []string `json:"synonyms,omitempty" yaml:"synonyms,omitempty"`

	// The value of the slot type. Must be less than or equal to 140 characters in length.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
