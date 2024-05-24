package types

type Lex_V2modelsSlotValueElicitationSettingPromptSpecification struct {
	//
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	//
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	//
	MessageGroups []Lex_V2modelsSlotValueElicitationSettingPromptSpecificationMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`

	//
	MessageSelectionStrategy string `json:"messageSelectionStrategy,omitempty" yaml:"messageSelectionStrategy,omitempty"`

	//
	PromptAttemptsSpecifications []Lex_V2modelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecification `json:"promptAttemptsSpecifications,omitempty" yaml:"promptAttemptsSpecifications,omitempty"`
}
