package types

type Lex_V2modelsIntentConfirmationSettingDeclinationNextStepDialogAction struct {
	// If the dialog action is `ElicitSlot`, defines the slot to elicit from the user.
	SlotToElicit string `json:"slotToElicit,omitempty" yaml:"slotToElicit,omitempty"`

	// Whether the next message for the intent is _not_ used.
	SuppressNextMessage bool `json:"suppressNextMessage,omitempty" yaml:"suppressNextMessage,omitempty"`

	// Action that the bot should execute. Valid values are `ElicitIntent`, `StartIntent`, `ElicitSlot`, `EvaluateConditional`, `InvokeDialogCodeHook`, `ConfirmIntent`, `FulfillIntent`, `CloseIntent`, `EndConversation`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
