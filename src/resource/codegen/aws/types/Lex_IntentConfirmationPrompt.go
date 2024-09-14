package types

type Lex_IntentConfirmationPrompt struct {
	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	MaxAttempts int `json:"maxAttempts,omitempty" yaml:"maxAttempts,omitempty"`

	//
	Messages []Lex_IntentConfirmationPromptMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
