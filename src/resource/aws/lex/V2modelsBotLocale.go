package lex

import types "DesignSphere_Server/src/resource/aws/types"

type V2modelsBotLocale struct {
	// Identifier of the bot to create the locale for.
	BotId string `json:"botId,omitempty" yaml:"botId,omitempty"`

	// Version of the bot to create the locale for. This can only be the draft version of the bot.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// Description of the bot locale. Use this to help identify the bot locale in lists.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Identifier of the language and locale that the bot will be used in. The string must match one of the supported locales. All of the intents, slot types, and slots used in the bot must have the same locale. For more information, see Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	LocaleId string `json:"localeId,omitempty" yaml:"localeId,omitempty"`

	/*
	   Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents.

	   The following arguments are optional:
	*/
	NLuIntentConfidenceThreshold float64 `json:"nLuIntentConfidenceThreshold,omitempty" yaml:"nLuIntentConfidenceThreshold,omitempty"`

	// Specified locale name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Timeouts types.Lex_V2modelsBotLocaleTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Amazon Polly voice ID that Amazon Lex uses for voice interaction with the user. See `voice_settings`.
	VoiceSettings types.Lex_V2modelsBotLocaleVoiceSettings `json:"voiceSettings,omitempty" yaml:"voiceSettings,omitempty"`
}
