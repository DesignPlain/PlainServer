package types

type Lex_V2modelsSlotValueElicitationSetting struct {
	//
	SampleUtterances []Lex_V2modelsSlotValueElicitationSettingSampleUtterance `json:"sampleUtterances,omitempty" yaml:"sampleUtterances,omitempty"`

	//
	SlotConstraint string `json:"slotConstraint,omitempty" yaml:"slotConstraint,omitempty"`

	//
	SlotResolutionSettings []Lex_V2modelsSlotValueElicitationSettingSlotResolutionSetting `json:"slotResolutionSettings,omitempty" yaml:"slotResolutionSettings,omitempty"`

	//
	WaitAndContinueSpecifications []Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecification `json:"waitAndContinueSpecifications,omitempty" yaml:"waitAndContinueSpecifications,omitempty"`

	//
	DefaultValueSpecifications []Lex_V2modelsSlotValueElicitationSettingDefaultValueSpecification `json:"defaultValueSpecifications,omitempty" yaml:"defaultValueSpecifications,omitempty"`

	//
	PromptSpecification Lex_V2modelsSlotValueElicitationSettingPromptSpecification `json:"promptSpecification,omitempty" yaml:"promptSpecification,omitempty"`
}
