package types

type Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariation struct {
	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationPlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`

	//
	CustomPayloads []Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationCustomPayload `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`
}
