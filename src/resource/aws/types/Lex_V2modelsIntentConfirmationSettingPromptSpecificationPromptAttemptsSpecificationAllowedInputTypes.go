package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes struct {
	// Whether audio input is allowed.
	AllowAudioInput bool `json:"allowAudioInput,omitempty" yaml:"allowAudioInput,omitempty"`

	// Whether DTMF input is allowed.
	AllowDtmfInput bool `json:"allowDtmfInput,omitempty" yaml:"allowDtmfInput,omitempty"`
}
