package types

type Medialive_ChannelEncoderSettingsGlobalConfigurationInputLossBehavior struct {
	//
	InputLossImageSlate Medialive_ChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate `json:"inputLossImageSlate,omitempty" yaml:"inputLossImageSlate,omitempty"`

	//
	InputLossImageType string `json:"inputLossImageType,omitempty" yaml:"inputLossImageType,omitempty"`

	//
	RepeatFrameMsec int `json:"repeatFrameMsec,omitempty" yaml:"repeatFrameMsec,omitempty"`

	//
	BlackFrameMsec int `json:"blackFrameMsec,omitempty" yaml:"blackFrameMsec,omitempty"`

	//
	InputLossImageColor string `json:"inputLossImageColor,omitempty" yaml:"inputLossImageColor,omitempty"`
}
