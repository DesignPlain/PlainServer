package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsCaptionLanguageMapping struct {
	//
	CaptionChannel int `json:"captionChannel,omitempty" yaml:"captionChannel,omitempty"`

	// Selects a specific three-letter language code from within an audio source.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	// Human readable information to indicate captions available for players (eg. English, or Spanish).
	LanguageDescription string `json:"languageDescription,omitempty" yaml:"languageDescription,omitempty"`
}
