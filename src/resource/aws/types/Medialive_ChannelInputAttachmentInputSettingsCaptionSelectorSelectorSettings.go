package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettings struct {
	// Ancillary Source Settings. See Ancillary Source Settings for more details.
	AncillarySourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAncillarySourceSettings `json:"ancillarySourceSettings,omitempty" yaml:"ancillarySourceSettings,omitempty"`

	// ARIB Source Settings.
	AribSourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsAribSourceSettings `json:"aribSourceSettings,omitempty" yaml:"aribSourceSettings,omitempty"`

	// DVB Sub Source Settings. See DVB Sub Source Settings for more details.
	DvbSubSourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsDvbSubSourceSettings `json:"dvbSubSourceSettings,omitempty" yaml:"dvbSubSourceSettings,omitempty"`

	// Embedded Source Settings. See Embedded Source Settings for more details.
	EmbeddedSourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsEmbeddedSourceSettings `json:"embeddedSourceSettings,omitempty" yaml:"embeddedSourceSettings,omitempty"`

	// SCTE20 Source Settings. See SCTE 20 Source Settings for more details.
	Scte20SourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsScte20SourceSettings `json:"scte20SourceSettings,omitempty" yaml:"scte20SourceSettings,omitempty"`

	// SCTE27 Source Settings. See SCTE 27 Source Settings for more details.
	Scte27SourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsScte27SourceSettings `json:"scte27SourceSettings,omitempty" yaml:"scte27SourceSettings,omitempty"`

	// Teletext Source Settings. See Teletext Source Settings for more details.
	TeletextSourceSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettings `json:"teletextSourceSettings,omitempty" yaml:"teletextSourceSettings,omitempty"`
}
