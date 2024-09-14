package types

type Lex_IntentSlot struct {
	// The name of the intent slot that you want to create. The name is case sensitive. Must be less than or equal to 100 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The response card. Amazon Lex will substitute session attributes and
	   slot values into the response card. For more information, see
	   [Example: Using a Response Card](https://docs.aws.amazon.com/lex/latest/dg/ex-resp-card.html). Must be less than or equal to 50000 characters in length.
	*/
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`

	/*
	   If you know a specific pattern with which users might respond to
	   an Amazon Lex request for a slot value, you can provide those utterances to improve accuracy. This
	   is optional. In most cases, Amazon Lex is capable of understanding user utterances. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	*/
	SampleUtterances []string `json:"sampleUtterances,omitempty" yaml:"sampleUtterances,omitempty"`

	// Specifies whether the slot is required or optional.
	SlotConstraint string `json:"slotConstraint,omitempty" yaml:"slotConstraint,omitempty"`

	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Directs Lex the order in which to elicit this slot value from the user.
	   For example, if the intent has two slots with priorities 1 and 2, AWS Lex first elicits a value for
	   the slot with priority 1. If multiple slots share the same priority, the order in which Lex elicits
	   values is arbitrary. Must be between 1 and 100.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The type of the slot, either a custom slot type that you defined or one of
	   the built-in slot types. Must be less than or equal to 100 characters in length.
	*/
	SlotType string `json:"slotType,omitempty" yaml:"slotType,omitempty"`

	// The version of the slot type. Must be less than or equal to 64 characters in length.
	SlotTypeVersion string `json:"slotTypeVersion,omitempty" yaml:"slotTypeVersion,omitempty"`

	/*
	   The prompt that Amazon Lex uses to elicit the slot value
	   from the user. Attributes are documented under prompt.
	*/
	ValueElicitationPrompt Lex_IntentSlotValueElicitationPrompt `json:"valueElicitationPrompt,omitempty" yaml:"valueElicitationPrompt,omitempty"`
}
