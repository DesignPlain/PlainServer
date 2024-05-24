package types

type Polly_getVoicesVoice struct {
	// Amazon Polly assigned voice ID.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Language identification tag for filtering the list of voices returned. If not specified, all available voices are returned.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// Human readable name of the language in English.
	LanguageName string `json:"languageName,omitempty" yaml:"languageName,omitempty"`

	// Name of the voice.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies which engines are supported by a given voice.
	SupportedEngines []string `json:"supportedEngines,omitempty" yaml:"supportedEngines,omitempty"`

	// Additional codes for languages available for the specified voice in addition to its default language.
	AdditionalLanguageCodes []string `json:"additionalLanguageCodes,omitempty" yaml:"additionalLanguageCodes,omitempty"`

	// Gender of the voice.
	Gender string `json:"gender,omitempty" yaml:"gender,omitempty"`
}
