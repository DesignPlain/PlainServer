package types

type Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariation struct {
	//
	PlainTextMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessage `json:"plainTextMessage,omitempty" yaml:"plainTextMessage,omitempty"`

	//
	SsmlMessage Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessage `json:"ssmlMessage,omitempty" yaml:"ssmlMessage,omitempty"`

	//
	CustomPayloads []string `json:"customPayloads,omitempty" yaml:"customPayloads,omitempty"`

	//
	ImageResponseCard Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCard `json:"imageResponseCard,omitempty" yaml:"imageResponseCard,omitempty"`
}
