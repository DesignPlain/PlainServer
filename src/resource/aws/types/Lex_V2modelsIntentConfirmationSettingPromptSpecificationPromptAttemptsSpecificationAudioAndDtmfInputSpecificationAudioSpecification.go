package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecification struct {
	// Time for which a bot waits after the customer stops speaking to assume the utterance is finished.
	EndTimeoutMs int `json:"endTimeoutMs,omitempty" yaml:"endTimeoutMs,omitempty"`

	// Time for how long Amazon Lex waits before speech input is truncated and the speech is returned to application.
	MaxLengthMs int `json:"maxLengthMs,omitempty" yaml:"maxLengthMs,omitempty"`
}
