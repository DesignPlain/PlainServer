package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettings struct {
	// Set the colorspace to Rec. 709.
	Rec709Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettingsRec709Settings `json:"rec709Settings,omitempty" yaml:"rec709Settings,omitempty"`

	// Sets the colorspace metadata to be passed through.
	ColorSpacePassthroughSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettingsColorSpacePassthroughSettings `json:"colorSpacePassthroughSettings,omitempty" yaml:"colorSpacePassthroughSettings,omitempty"`

	// Set the colorspace to Dolby Vision81.
	DolbyVision81Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettingsDolbyVision81Settings `json:"dolbyVision81Settings,omitempty" yaml:"dolbyVision81Settings,omitempty"`

	// Set the colorspace to be HDR10. See H265 HDR10 Settings for more details.
	Hdr10Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings `json:"hdr10Settings,omitempty" yaml:"hdr10Settings,omitempty"`

	// Set the colorspace to Rec. 601.
	Rec601Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265SettingsColorSpaceSettingsRec601Settings `json:"rec601Settings,omitempty" yaml:"rec601Settings,omitempty"`
}
