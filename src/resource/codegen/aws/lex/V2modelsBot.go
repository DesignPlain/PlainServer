package lex

import types "libds/aws/types"

type V2modelsBot struct {
	// List of bot members in a network to be created. See `bot_members`.
	Members []types.Lex_V2modelsBotMember `json:"members,omitempty" yaml:"members,omitempty"`

	/*
	   ARN of an IAM role that has permission to access the bot.

	   The following arguments are optional:
	*/
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	//
	Timeouts types.Lex_V2modelsBotTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Type of a bot to create. Possible values are `"Bot"` and `"BotNetwork"`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Provides information on additional privacy protections Amazon Lex should use with the bot's data. See `data_privacy`
	DataPrivacies []types.Lex_V2modelsBotDataPrivacy `json:"dataPrivacies,omitempty" yaml:"dataPrivacies,omitempty"`

	// Description of the bot. It appears in lists to help you identify a particular bot.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Time, in seconds, that Amazon Lex should keep information about a user's conversation with the bot. You can specify between 60 (1 minute) and 86,400 (24 hours) seconds.
	IdleSessionTtlInSeconds int `json:"idleSessionTtlInSeconds,omitempty" yaml:"idleSessionTtlInSeconds,omitempty"`

	// Name of the bot. The bot name must be unique in the account that creates the bot. Type String. Length Constraints: Minimum length of 1. Maximum length of 100.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of tags to add to the bot. You can only add tags when you create a bot.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// List of tags to add to the test alias for a bot. You can only add tags when you create a bot.
	TestBotAliasTags map[string]string `json:"testBotAliasTags,omitempty" yaml:"testBotAliasTags,omitempty"`
}
