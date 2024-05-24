package types

type Lex_IntentFollowUpPrompt struct {
	// Prompts for information from the user. Attributes are documented under prompt.
	Prompt Lex_IntentFollowUpPromptPrompt `json:"prompt,omitempty" yaml:"prompt,omitempty"`

	/*
	   If the user answers "no" to the question defined in the prompt field,
	   Amazon Lex responds with this statement to acknowledge that the intent was canceled. Attributes are
	   documented below under statement.
	*/
	RejectionStatement Lex_IntentFollowUpPromptRejectionStatement `json:"rejectionStatement,omitempty" yaml:"rejectionStatement,omitempty"`
}
