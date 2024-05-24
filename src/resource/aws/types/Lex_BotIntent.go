package types

type Lex_BotIntent struct {
	// The name of the intent. Must be less than or equal to 100 characters in length.
	IntentName string `json:"intentName,omitempty" yaml:"intentName,omitempty"`

	// The version of the intent. Must be less than or equal to 64 characters in length.
	IntentVersion string `json:"intentVersion,omitempty" yaml:"intentVersion,omitempty"`
}
