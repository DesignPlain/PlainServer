package types

type Lex_V2modelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecification struct {
	//
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	//
	TextInputSpecification Lex_V2modelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification `json:"textInputSpecification,omitempty" yaml:"textInputSpecification,omitempty"`

	//
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	//
	AllowedInputTypes Lex_V2modelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes `json:"allowedInputTypes,omitempty" yaml:"allowedInputTypes,omitempty"`

	//
	AudioAndDtmfInputSpecification Lex_V2modelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification `json:"audioAndDtmfInputSpecification,omitempty" yaml:"audioAndDtmfInputSpecification,omitempty"`
}
