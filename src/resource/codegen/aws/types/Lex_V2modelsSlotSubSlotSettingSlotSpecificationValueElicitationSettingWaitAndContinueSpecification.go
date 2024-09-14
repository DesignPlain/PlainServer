package types

type Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification struct {
	//
	StillWaitingResponses []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponse `json:"stillWaitingResponses,omitempty" yaml:"stillWaitingResponses,omitempty"`

	//
	WaitingResponses []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationWaitingResponse `json:"waitingResponses,omitempty" yaml:"waitingResponses,omitempty"`

	//
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	//
	ContinueResponses []Lex_V2modelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationContinueResponse `json:"continueResponses,omitempty" yaml:"continueResponses,omitempty"`
}
