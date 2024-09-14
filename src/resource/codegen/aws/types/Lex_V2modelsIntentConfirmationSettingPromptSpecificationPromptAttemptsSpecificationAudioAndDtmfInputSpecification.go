package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification struct {
	// Configuration block for the settings on DTMF input. See `dtmf_specification`.
	DtmfSpecification Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification `json:"dtmfSpecification,omitempty" yaml:"dtmfSpecification,omitempty"`

	// Time for which a bot waits before assuming that the customer isn't going to speak or press a key. This timeout is shared between Audio and DTMF inputs.
	StartTimeoutMs int `json:"startTimeoutMs,omitempty" yaml:"startTimeoutMs,omitempty"`

	// Configuration block for the settings on audio input. See `audio_specification`.
	AudioSpecification Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecification `json:"audioSpecification,omitempty" yaml:"audioSpecification,omitempty"`
}
