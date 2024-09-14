package types

type Lex_BotClarificationPrompt struct {
	// The number of times to prompt the user for information.
	MaxAttempts int `json:"maxAttempts,omitempty" yaml:"maxAttempts,omitempty"`

	//
	Messages []Lex_BotClarificationPromptMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
