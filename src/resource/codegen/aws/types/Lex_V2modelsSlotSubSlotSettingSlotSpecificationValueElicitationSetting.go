package types

type Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting struct {
	//
	WaitAndContinueSpecifications []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification `json:"waitAndContinueSpecifications,omitempty" yaml:"waitAndContinueSpecifications,omitempty"`

	//
	DefaultValueSpecifications []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification `json:"defaultValueSpecifications,omitempty" yaml:"defaultValueSpecifications,omitempty"`

	//
	PromptSpecification Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification `json:"promptSpecification,omitempty" yaml:"promptSpecification,omitempty"`

	//
	SampleUtterances []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance `json:"sampleUtterances,omitempty" yaml:"sampleUtterances,omitempty"`
}
