package types

type Lex_BotAbortStatement struct {
	//
	Messages []Lex_BotAbortStatementMessage `json:"messages,omitempty" yaml:"messages,omitempty"`

	//
	ResponseCard string `json:"responseCard,omitempty" yaml:"responseCard,omitempty"`
}
