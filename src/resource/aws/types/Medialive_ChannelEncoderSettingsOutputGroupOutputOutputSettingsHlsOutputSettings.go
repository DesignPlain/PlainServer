package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettings struct {
	//
	SegmentModifier string `json:"segmentModifier,omitempty" yaml:"segmentModifier,omitempty"`

	//
	H265PackagingType string `json:"h265PackagingType,omitempty" yaml:"h265PackagingType,omitempty"`

	//
	HlsSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsHlsOutputSettingsHlsSettings `json:"hlsSettings,omitempty" yaml:"hlsSettings,omitempty"`

	// String concatenated to the end of the destination filename. Required for multiple outputs of the same type.
	NameModifier string `json:"nameModifier,omitempty" yaml:"nameModifier,omitempty"`
}
