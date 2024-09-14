package types

type Lex_IntentRejectionStatement struct {
	//
	Messages []Lex_IntentRejectionStatementMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
