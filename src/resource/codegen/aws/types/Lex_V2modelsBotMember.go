package types

type Lex_V2modelsBotMember struct {
	// (Required) - Unique ID of a bot that is a member of this network of bots.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the bot. The bot name must be unique in the account that creates the bot. Type String. Length Constraints: Minimum length of 1. Maximum length of 100.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// (Required) - Version of a bot that is a member of this network of bots.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// (Required) - Alias ID of a bot that is a member of this network of bots.
	AliasId string `json:"aliasId,omitempty" yaml:"aliasId,omitempty"`

	// (Required) - Alias name of a bot that is a member of this network of bots.
	AliasName string `json:"aliasName,omitempty" yaml:"aliasName,omitempty"`
}
