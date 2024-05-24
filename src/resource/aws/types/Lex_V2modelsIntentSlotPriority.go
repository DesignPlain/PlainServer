package types

type Lex_V2modelsIntentSlotPriority struct {
	// Unique identifier of the slot.
	SlotId string `json:"slotId,omitempty" yaml:"slotId,omitempty"`

	// Priority that Amazon Lex should apply to the slot.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
