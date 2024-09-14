package types

type Lex_V2modelsBotLocaleVoiceSettings struct {
	// Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user. Valid values are `standard` and `neural`. If not specified, the default is `standard`.
	Engine string `json:"engine,omitempty" yaml:"engine,omitempty"`

	// Identifier of the Amazon Polly voice to use.
	VoiceId string `json:"voiceId,omitempty" yaml:"voiceId,omitempty"`
}
