package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecification struct {
	// Whether the user can interrupt a speech prompt attempt from the bot.
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	// Configuration block for the allowed input types of the prompt attempt. See `allowed_input_types`.
	AllowedInputTypes Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes `json:"allowedInputTypes,omitempty" yaml:"allowedInputTypes,omitempty"`

	// Configuration block for settings on audio and DTMF input. See `audio_and_dtmf_input_specification`.
	AudioAndDtmfInputSpecification Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification `json:"audioAndDtmfInputSpecification,omitempty" yaml:"audioAndDtmfInputSpecification,omitempty"`

	// Which attempt to configure. Valid values are `Initial`, `Retry1`, `Retry2`, `Retry3`, `Retry4`, `Retry5`.
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	// Configuration block for the settings on text input. See `text_input_specification`.
	TextInputSpecification Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification `json:"textInputSpecification,omitempty" yaml:"textInputSpecification,omitempty"`
}
