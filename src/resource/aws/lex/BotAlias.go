package lex

import types "DesignSphere_Server/src/resource/aws/types"

type BotAlias struct {
	// The name of the bot.
	BotName string `json:"botName,omitempty" yaml:"botName,omitempty"`

	// The version of the bot.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
	ConversationLogs types.Lex_BotAliasConversationLogs `json:"conversationLogs,omitempty" yaml:"conversationLogs,omitempty"`

	// A description of the alias. Must be less than or equal to 200 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the alias. The name is not case sensitive. Must be less than or equal to 100 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
