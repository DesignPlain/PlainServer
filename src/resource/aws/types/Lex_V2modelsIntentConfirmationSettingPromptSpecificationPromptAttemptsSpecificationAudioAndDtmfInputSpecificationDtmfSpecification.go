package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification struct {
	// DTMF character that clears the accumulated DTMF digits and immediately ends the input.
	DeletionCharacter string `json:"deletionCharacter,omitempty" yaml:"deletionCharacter,omitempty"`

	// DTMF character that immediately ends input. If the user does not press this character, the input ends after the end timeout.
	EndCharacter string `json:"endCharacter,omitempty" yaml:"endCharacter,omitempty"`

	// How long the bot should wait after the last DTMF character input before assuming that the input has concluded.
	EndTimeoutMs int `json:"endTimeoutMs,omitempty" yaml:"endTimeoutMs,omitempty"`

	// Maximum number of DTMF digits allowed in an utterance.
	MaxLength int `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
}
