package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettingsAudioOnlyHlsSettings struct {
	//
	AudioTrackType string `json:"audioTrackType,omitempty" yaml:"audioTrackType,omitempty"`

	//
	SegmentType string `json:"segmentType,omitempty" yaml:"segmentType,omitempty"`

	//
	AudioGroupId string `json:"audioGroupId,omitempty" yaml:"audioGroupId,omitempty"`

	//
	AudioOnlyImage Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettingsAudioOnlyHlsSettingsAudioOnlyImage `json:"audioOnlyImage,omitempty" yaml:"audioOnlyImage,omitempty"`
}
