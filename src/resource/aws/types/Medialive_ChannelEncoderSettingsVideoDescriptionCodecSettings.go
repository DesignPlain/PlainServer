package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettings struct {
	//
	H265Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH265Settings `json:"h265Settings,omitempty" yaml:"h265Settings,omitempty"`

	// Frame capture settings. See Frame Capture Settings for more details.
	FrameCaptureSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettings `json:"frameCaptureSettings,omitempty" yaml:"frameCaptureSettings,omitempty"`

	// H264 settings. See H264 Settings for more details.
	H264Settings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsH264Settings `json:"h264Settings,omitempty" yaml:"h264Settings,omitempty"`
}
