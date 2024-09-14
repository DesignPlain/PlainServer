package types

type Lex_IntentConclusionStatement struct {
	//
	Messages []Lex_IntentConclusionStatementMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
