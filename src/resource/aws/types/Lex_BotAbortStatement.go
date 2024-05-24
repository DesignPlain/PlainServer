package types

type Lex_BotAbortStatement struct {
	/*
	   A set of messages, each of which provides a message string and its type.
	   You can specify the message string in plain text or in Speech Synthesis Markup Language (SSML).
	   Attributes are documented under message.
	*/
	Messages []Lex_BotAbortStatementMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	/*
	   The response card. Amazon Lex will substitute session attributes and
	   slot values into the response card. For more information, see
	   [Example: Using a Response Card](https://docs.aws.amazon.com/lex/latest/dg/ex-resp-card.html).
	*/
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
