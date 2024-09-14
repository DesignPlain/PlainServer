package types

type Medialive_ChannelEncoderSettingsCaptionDescription struct {
	// Human readable information to indicate captions available for players (eg. English, or Spanish).
	LanguageDescription string `json:"languageDescription,omitempty" yaml:"languageDescription,omitempty"`

	// Name of the caption description. Used to associate a caption description with an output. Names must be unique within an event.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Indicates whether the caption track implements accessibility features such as written descriptions of spoken dialog, music, and sounds.
	Accessibility string `json:"accessibility,omitempty" yaml:"accessibility,omitempty"`

	// Specifies which input caption selector to use as a caption source when generating output captions. This field should match a captionSelector name.
	CaptionSelectorName string `json:"captionSelectorName,omitempty" yaml:"captionSelectorName,omitempty"`

	// Additional settings for captions destination that depend on the destination type. See Destination Settings for more details.
	DestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettings `json:"destinationSettings,omitempty" yaml:"destinationSettings,omitempty"`

	// ISO 639-2 three-digit code.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`
}
