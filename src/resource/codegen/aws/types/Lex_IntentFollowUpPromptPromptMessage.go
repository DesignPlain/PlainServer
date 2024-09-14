package types

type Lex_IntentFollowUpPromptPromptMessage struct {
	// The content type of the message string.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	/*
	   Identifies the message group that the message belongs to. When a group
	   is assigned to a message, Amazon Lex returns one message from each group in the response. Must be a number between 1 and 5 (inclusive).
	*/
	GroupNumber int `json:"groupNumber,omitempty" yaml:"groupNumber,omitempty"`

	// The text of the message. Must be less than or equal to 1000 characters in length.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`
}
