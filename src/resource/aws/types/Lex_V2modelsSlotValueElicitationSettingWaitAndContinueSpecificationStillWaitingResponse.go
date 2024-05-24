package types

type Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponse struct {
	//
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	//
	FrequencyInSeconds int `json:"frequencyInSeconds,omitempty" yaml:"frequencyInSeconds,omitempty"`

	//
	MessageGroups []Lex_V2modelsSlotValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`

	//
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" yaml:"timeoutInSeconds,omitempty"`
}
