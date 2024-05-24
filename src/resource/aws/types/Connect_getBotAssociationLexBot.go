package types

type Connect_getBotAssociationLexBot struct {
	// Region that the Amazon Lex (V1) bot was created in.
	LexRegion string `json:"lexRegion,omitempty" yaml:"lexRegion,omitempty"`

	// Name of the Amazon Lex (V1) bot.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
