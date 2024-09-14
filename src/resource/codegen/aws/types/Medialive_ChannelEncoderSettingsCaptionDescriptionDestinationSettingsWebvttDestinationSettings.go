package types

type Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsWebvttDestinationSettings struct {
	// Controls whether the color and position of the source captions is passed through to the WebVTT output captions. PASSTHROUGH - Valid only if the source captions are EMBEDDED or TELETEXT. NO\_STYLE\_DATA - Don’t pass through the style. The output captions will not contain any font styling information.
	StyleControl string `json:"styleControl,omitempty" yaml:"styleControl,omitempty"`
}
