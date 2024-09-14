package types

type Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecification struct {
	//
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	//
	AllowedInputTypes Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes `json:"allowedInputTypes,omitempty" yaml:"allowedInputTypes,omitempty"`

	//
	AudioAndDtmfInputSpecification Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification `json:"audioAndDtmfInputSpecification,omitempty" yaml:"audioAndDtmfInputSpecification,omitempty"`

	//
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	//
	TextInputSpecification Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification `json:"textInputSpecification,omitempty" yaml:"textInputSpecification,omitempty"`
}
