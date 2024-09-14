package types

type Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification struct {
	//
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	//
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	//
	MessageGroups []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`

	//
	MessageSelectionStrategy string `json:"messageSelectionStrategy,omitempty" yaml:"messageSelectionStrategy,omitempty"`

	//
	PromptAttemptsSpecifications []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecification `json:"promptAttemptsSpecifications,omitempty" yaml:"promptAttemptsSpecifications,omitempty"`
}
