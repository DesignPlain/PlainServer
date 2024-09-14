package types

type Lex_V2modelsIntentConfirmationSettingPromptSpecification struct {
	// Whether the user can interrupt a speech prompt from the bot.
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	// Maximum number of times the bot tries to elicit a response from the user using this prompt.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`

	// Configuration block for messages that Amazon Lex can send to the user. Amazon Lex chooses the actual message to send at runtime. See `message_group`.
	MessageGroups []Lex_V2modelsIntentConfirmationSettingPromptSpecificationMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`

	// How a message is selected from a message group among retries. Valid values are `Random` and `Ordered`.
	MessageSelectionStrategy string `json:"messageSelectionStrategy,omitempty" yaml:"messageSelectionStrategy,omitempty"`

	// Configuration block for advanced settings on each attempt of the prompt. See `prompt_attempts_specification`.
	PromptAttemptsSpecifications []Lex_V2modelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecification `json:"promptAttemptsSpecifications,omitempty" yaml:"promptAttemptsSpecifications,omitempty"`
}
