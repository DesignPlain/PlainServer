package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettings struct {
	//
	MultiplexGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMultiplexGroupSettings `json:"multiplexGroupSettings,omitempty" yaml:"multiplexGroupSettings,omitempty"`

	// RTMP group settings. See RTMP Group Settings for more details.
	RtmpGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettings `json:"rtmpGroupSettings,omitempty" yaml:"rtmpGroupSettings,omitempty"`

	//
	UdpGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsUdpGroupSettings `json:"udpGroupSettings,omitempty" yaml:"udpGroupSettings,omitempty"`

	// Archive group settings. See Archive Group Settings for more details.
	ArchiveGroupSettings []Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSetting `json:"archiveGroupSettings,omitempty" yaml:"archiveGroupSettings,omitempty"`

	//
	FrameCaptureGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettings `json:"frameCaptureGroupSettings,omitempty" yaml:"frameCaptureGroupSettings,omitempty"`

	//
	HlsGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettings `json:"hlsGroupSettings,omitempty" yaml:"hlsGroupSettings,omitempty"`

	// Media package group settings. See Media Package Group Settings for more details.
	MediaPackageGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMediaPackageGroupSettings `json:"mediaPackageGroupSettings,omitempty" yaml:"mediaPackageGroupSettings,omitempty"`

	//
	MsSmoothGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettings `json:"msSmoothGroupSettings,omitempty" yaml:"msSmoothGroupSettings,omitempty"`
}
