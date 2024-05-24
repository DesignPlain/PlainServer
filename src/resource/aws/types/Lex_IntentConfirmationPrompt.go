package types

type Lex_IntentConfirmationPrompt struct {
	// The number of times to prompt the user for information. Must be a number between 1 and 5 (inclusive).
	MaxAttempts int `json:"maxAttempts,omitempty" yaml:"maxAttempts,omitempty"`

	/*
	   A set of messages, each of which provides a message string and its type.
	   You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	   Attributes are documented under message. Must contain between 1 and 15 messages.
	*/
	Messages []Lex_IntentConfirmationPromptMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	/*
	   The response card. Amazon Lex will substitute session attributes and
	   slot values into the response card. For more information, see
	   [Example: Using a Response Card](https://docs.aws.amazon.com/lex/latest/dg/ex-resp-card.html). Must be less than or equal to 50000 characters in length.
	*/
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
