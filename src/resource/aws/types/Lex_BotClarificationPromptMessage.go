package types

type Lex_BotClarificationPromptMessage struct {
	/*
	   Identifies the message group that the message belongs to. When a group
	   is assigned to a message, Amazon Lex returns one message from each group in the response.
	*/
	GroupNumber int `json:"groupNumber,omitempty" yaml:"groupNumber,omitempty"`

	// The text of the message.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The content type of the message string.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
}
