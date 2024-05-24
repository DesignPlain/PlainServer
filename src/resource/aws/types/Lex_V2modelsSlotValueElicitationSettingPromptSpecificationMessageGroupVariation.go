package types

type Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariation struct {
	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`

	//
	CustomPayloads []string `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`

	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`

	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroupVariationPlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`
}
