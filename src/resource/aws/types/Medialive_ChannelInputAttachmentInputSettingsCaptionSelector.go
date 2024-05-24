package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelector struct {
	// Selects a specific three-letter language code from within an audio source.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   Name of the Channel.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The audio selector settings. See Audio Selector Settings for more details.
	SelectorSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettings `json:"selectorSettings,omitempty" yaml:"selectorSettings,omitempty"`
}
