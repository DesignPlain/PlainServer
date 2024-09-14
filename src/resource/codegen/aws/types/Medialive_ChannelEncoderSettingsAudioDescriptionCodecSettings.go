package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettings struct {
	//
	Mp2Settings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsMp2Settings `json:"mp2Settings,omitempty" yaml:"mp2Settings,omitempty"`

	//
	PassThroughSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsPassThroughSettings `json:"passThroughSettings,omitempty" yaml:"passThroughSettings,omitempty"`

	//
	WavSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettings `json:"wavSettings,omitempty" yaml:"wavSettings,omitempty"`

	// Aac Settings. See AAC Settings for more details.
	AacSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettings `json:"aacSettings,omitempty" yaml:"aacSettings,omitempty"`

	// Ac3 Settings. See AC3 Settings for more details.
	Ac3Settings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsAc3Settings `json:"ac3Settings,omitempty" yaml:"ac3Settings,omitempty"`

	// Eac3 Atmos Settings. See EAC3 Atmos Settings
	Eac3AtmosSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsEac3AtmosSettings `json:"eac3AtmosSettings,omitempty" yaml:"eac3AtmosSettings,omitempty"`

	// Eac3 Settings. See EAC3 Settings
	Eac3Settings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsEac3Settings `json:"eac3Settings,omitempty" yaml:"eac3Settings,omitempty"`
}
