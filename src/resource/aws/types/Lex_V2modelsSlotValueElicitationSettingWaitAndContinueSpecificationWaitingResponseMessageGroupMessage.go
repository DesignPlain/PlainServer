package types

type Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupMessage struct {
	//
	CustomPayloads []string `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`

	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupMessageImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupMessagePlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponseMessageGroupMessageSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`
}
