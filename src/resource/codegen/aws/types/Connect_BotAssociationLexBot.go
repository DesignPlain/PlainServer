package types

type Connect_BotAssociationLexBot struct {
	// The Region that the Amazon Lex (V1) bot was created in. Defaults to current region.
	LexRegion string `json:"lexRegion,omitempty" yaml:"lexRegion,omitempty"`

	// The name of the Amazon Lex (V1) bot.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
