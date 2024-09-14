package types

type Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupMessage struct {
	//
	CustomPayloads []Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupMessageCustomPayload `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`

	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupMessageImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupMessagePlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupMessageSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`
}
