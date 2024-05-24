package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsTimecodeBurninSettings struct {
	// Set a prefix on the burned in timecode.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Sets the size of the burned in timecode.
	TimecodeBurninFontSize string `json:"timecodeBurninFontSize,omitempty" yaml:"timecodeBurninFontSize,omitempty"`

	// Sets the position of the burned in timecode.
	TimecodeBurninPosition string `json:"timecodeBurninPosition,omitempty" yaml:"timecodeBurninPosition,omitempty"`
}
