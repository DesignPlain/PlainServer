package types

type Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecification struct {
	//
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	//
	ContinueResponses []Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationContinueResponse `json:"continueResponses,omitempty" yaml:"continueResponses,omitempty"`

	//
	StillWaitingResponses []Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponse `json:"stillWaitingResponses,omitempty" yaml:"stillWaitingResponses,omitempty"`

	//
	WaitingResponses []Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationWaitingResponse `json:"waitingResponses,omitempty" yaml:"waitingResponses,omitempty"`
}
