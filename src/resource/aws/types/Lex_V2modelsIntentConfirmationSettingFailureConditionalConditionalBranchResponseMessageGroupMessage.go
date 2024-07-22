package types

type Lex_V2modelsIntentConfirmationSettingFailureConditionalConditionalBranchResponseMessageGroupMessage struct {
	// Configuration block for a message in a custom format defined by the client application. See `custom_payload`.
	CustomPayload Lex_V2modelsIntentConfirmationSettingFailureConditionalConditionalBranchResponseMessageGroupMessageCustomPayload `json:"customPayload,omitempty" yaml:"customPayload,omitempty"`

	// Configuration block for a message that defines a response card that the client application can show to the user. See `image_response_card`.
	ImageResponseCard Lex_V2modelsIntentConfirmationSettingFailureConditionalConditionalBranchResponseMessageGroupMessageImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	// Configuration block for a message in plain text format. See `plain_text_message`.
	PlainTextMessage Lex_V2modelsIntentConfirmationSettingFailureConditionalConditionalBranchResponseMessageGroupMessagePlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	// Configuration block for a message in Speech Synthesis Markup Language (SSML). See `ssml_message`.
	SsmlMessage Lex_V2modelsIntentConfirmationSettingFailureConditionalConditionalBranchResponseMessageGroupMessageSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`
}