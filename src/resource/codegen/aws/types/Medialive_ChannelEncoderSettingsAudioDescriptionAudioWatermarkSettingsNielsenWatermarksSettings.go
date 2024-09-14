package types

type Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettings struct {
	// Used to insert watermarks of type Nielsen CBET. See Nielsen CBET Settings for more details.
	NielsenCbetSettings Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenCbetSettings `json:"nielsenCbetSettings,omitempty" yaml:"nielsenCbetSettings,omitempty"`

	// Distribution types to assign to the watermarks. Options are `PROGRAM_CONTENT` and `FINAL_DISTRIBUTOR`.
	NielsenDistributionType string `json:"nielsenDistributionType,omitempty" yaml:"nielsenDistributionType,omitempty"`

	// Used to insert watermarks of type Nielsen NAES, II (N2) and Nielsen NAES VI (NW). See Nielsen NAES II NW Settings for more details.
	NielsenNaesIiNwSettings []Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettingsNielsenWatermarksSettingsNielsenNaesIiNwSetting `json:"nielsenNaesIiNwSettings,omitempty" yaml:"nielsenNaesIiNwSettings,omitempty"`
}
