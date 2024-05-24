package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettings struct {
	// M2TS Settings. See [M2TS Settings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m2tssettings.html) for more details.
	M2tsSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettings `json:"m2tsSettings,omitempty" yaml:"m2tsSettings,omitempty"`

	// Raw Settings. This can be set as an empty block.
	RawSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsRawSettings `json:"rawSettings,omitempty" yaml:"rawSettings,omitempty"`
}
