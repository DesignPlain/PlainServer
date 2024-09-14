package types

type Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettings struct {
	// Teletext Destination Settings.
	TeletextDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTeletextDestinationSettings `json:"teletextDestinationSettings,omitempty" yaml:"teletextDestinationSettings,omitempty"`

	// WebVTT Destination Settings. See WebVTT Destination Settings for more details.
	WebvttDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsWebvttDestinationSettings `json:"webvttDestinationSettings,omitempty" yaml:"webvttDestinationSettings,omitempty"`

	// Burn In Destination Settings. See Burn In Destination Settings for more details.
	BurnInDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsBurnInDestinationSettings `json:"burnInDestinationSettings,omitempty" yaml:"burnInDestinationSettings,omitempty"`

	// DVB Sub Destination Settings. See DVB Sub Destination Settings for more details.
	DvbSubDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsDvbSubDestinationSettings `json:"dvbSubDestinationSettings,omitempty" yaml:"dvbSubDestinationSettings,omitempty"`

	// EBU TT D Destination Settings. See EBU TT D Destination Settings for more details.
	EbuTtDDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsEbuTtDDestinationSettings `json:"ebuTtDDestinationSettings,omitempty" yaml:"ebuTtDDestinationSettings,omitempty"`

	// RTMP Caption Info Destination Settings.
	RtmpCaptionInfoDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsRtmpCaptionInfoDestinationSettings `json:"rtmpCaptionInfoDestinationSettings,omitempty" yaml:"rtmpCaptionInfoDestinationSettings,omitempty"`

	// SMPTE TT Destination Settings.
	SmpteTtDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsSmpteTtDestinationSettings `json:"smpteTtDestinationSettings,omitempty" yaml:"smpteTtDestinationSettings,omitempty"`

	// TTML Destination Settings. See TTML Destination Settings for more details.
	TtmlDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsTtmlDestinationSettings `json:"ttmlDestinationSettings,omitempty" yaml:"ttmlDestinationSettings,omitempty"`

	// ARIB Destination Settings.
	AribDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsAribDestinationSettings `json:"aribDestinationSettings,omitempty" yaml:"aribDestinationSettings,omitempty"`

	// Embedded Destination Settings.
	EmbeddedDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsEmbeddedDestinationSettings `json:"embeddedDestinationSettings,omitempty" yaml:"embeddedDestinationSettings,omitempty"`

	// Embedded Plus SCTE20 Destination Settings.
	EmbeddedPlusScte20DestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsEmbeddedPlusScte20DestinationSettings `json:"embeddedPlusScte20DestinationSettings,omitempty" yaml:"embeddedPlusScte20DestinationSettings,omitempty"`

	// SCTE20 Plus Embedded Destination Settings.
	Scte20PlusEmbeddedDestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte20PlusEmbeddedDestinationSettings `json:"scte20PlusEmbeddedDestinationSettings,omitempty" yaml:"scte20PlusEmbeddedDestinationSettings,omitempty"`

	// SCTE27 Destination Settings.
	Scte27DestinationSettings Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsScte27DestinationSettings `json:"scte27DestinationSettings,omitempty" yaml:"scte27DestinationSettings,omitempty"`
}
