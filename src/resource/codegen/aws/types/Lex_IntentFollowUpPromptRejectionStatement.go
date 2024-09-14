package types

type Lex_IntentFollowUpPromptRejectionStatement struct {
	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`

	//
	Messages []Lex_IntentFollowUpPromptRejectionStatementMessage `json:"messages,omitempty" yaml:"messages,omitempty"`
}
