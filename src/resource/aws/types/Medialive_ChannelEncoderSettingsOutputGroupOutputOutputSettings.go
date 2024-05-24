package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettings struct {
	// Multiplex output settings. See Multiplex Output Settings for more details.
	MultiplexOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsMultiplexOutputSettings `json:"multiplexOutputSettings,omitempty" yaml:"multiplexOutputSettings,omitempty"`

	// RTMP output settings. See RTMP Output Settings for more details.
	RtmpOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettings `json:"rtmpOutputSettings,omitempty" yaml:"rtmpOutputSettings,omitempty"`

	// UDP output settings. See UDP Output Settings for more details.
	UdpOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettings `json:"udpOutputSettings,omitempty" yaml:"udpOutputSettings,omitempty"`

	// Archive output settings. See Archive Output Settings for more details.
	ArchiveOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettings `json:"archiveOutputSettings,omitempty" yaml:"archiveOutputSettings,omitempty"`

	//
	FrameCaptureOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsFrameCaptureOutputSettings `json:"frameCaptureOutputSettings,omitempty" yaml:"frameCaptureOutputSettings,omitempty"`

	//
	HlsOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings `json:"hlsOutputSettings,omitempty" yaml:"hlsOutputSettings,omitempty"`

	// Media package output settings. This can be set as an empty block.
	MediaPackageOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsMediaPackageOutputSettings `json:"mediaPackageOutputSettings,omitempty" yaml:"mediaPackageOutputSettings,omitempty"`

	//
	MsSmoothOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsMsSmoothOutputSettings `json:"msSmoothOutputSettings,omitempty" yaml:"msSmoothOutputSettings,omitempty"`
}
