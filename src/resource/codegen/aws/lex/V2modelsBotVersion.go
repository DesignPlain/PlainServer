package lex

import types "libds/aws/types"

type V2modelsBotVersion struct {
	// Idientifier of the bot to create the version for.
	BotId string `json:"botId,omitempty" yaml:"botId,omitempty"`

	// Version number assigned to the version.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// A description of the version. Use the description to help identify the version in lists.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specifies the locales that Amazon Lex adds to this version. You can choose the draft version or any other previously published version for each locale. When you specify a source version, the locale data is copied from the source version to the new version.

	   The attribute value is a map with one or more entries, each of which has a locale name as the key and an object with the following attribute as the value:
	   - `sourceBotVersion` - (Required) The version of a bot used for a bot locale. Valid values: `DRAFT`, a numeric version.
	*/
	LocaleSpecification map[string]types.Lex_V2modelsBotVersionLocaleSpecification `json:"localeSpecification,omitempty" yaml:"localeSpecification,omitempty"`

	//
	Timeouts types.Lex_V2modelsBotVersionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
