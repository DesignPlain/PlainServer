package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings struct {
	// Specifies the 608/708 channel number within the video track from which to extract captions. Unused for passthrough.
	Source608ChannelNumber int `json:"source608ChannelNumber,omitempty" yaml:"source608ChannelNumber,omitempty"`

	// If upconvert, 608 data is both passed through via the “608 compatibility bytes” fields of the 708 wrapper as well as translated into 708. 708 data present in the source content will be discarded.
	Convert608To708 string `json:"convert608To708,omitempty" yaml:"convert608To708,omitempty"`

	// Set to “auto” to handle streams with intermittent and/or non-aligned SCTE-20 and Embedded captions.
	Scte20Detection string `json:"scte20Detection,omitempty" yaml:"scte20Detection,omitempty"`
}
