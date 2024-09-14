package lex

import types "libds/aws/types"

type Bot struct {
	// Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enable_model_improvements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
	NluIntentConfidenceThreshold float64 `json:"nluIntentConfidenceThreshold,omitempty" yaml:"nluIntentConfidenceThreshold,omitempty"`

	// A description of the bot. Must be less than or equal to 200 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
	IdleSessionTtlInSeconds int `json:"idleSessionTtlInSeconds,omitempty" yaml:"idleSessionTtlInSeconds,omitempty"`

	// The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If you set the `process_behavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
	ProcessBehavior string `json:"processBehavior,omitempty" yaml:"processBehavior,omitempty"`

	// By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
	ChildDirected bool `json:"childDirected,omitempty" yaml:"childDirected,omitempty"`

	// The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
	ClarificationPrompt types.Lex_BotClarificationPrompt `json:"clarificationPrompt,omitempty" yaml:"clarificationPrompt,omitempty"`

	// A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 250 Intent objects.
	Intents []types.Lex_BotIntent `json:"intents,omitempty" yaml:"intents,omitempty"`

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
	VoiceId string `json:"voiceId,omitempty" yaml:"voiceId,omitempty"`

	// Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
	CreateVersion bool `json:"createVersion,omitempty" yaml:"createVersion,omitempty"`

	// Set to `true` to enable access to natural language understanding improvements. When you set the `enable_model_improvements` parameter to true you can use the `nlu_intent_confidence_threshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enable_model_improvements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
	EnableModelImprovements bool `json:"enableModelImprovements,omitempty" yaml:"enableModelImprovements,omitempty"`

	// Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
	Locale string `json:"locale,omitempty" yaml:"locale,omitempty"`

	// The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
	AbortStatement types.Lex_BotAbortStatement `json:"abortStatement,omitempty" yaml:"abortStatement,omitempty"`

	// When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
	DetectSentiment bool `json:"detectSentiment,omitempty" yaml:"detectSentiment,omitempty"`
}
