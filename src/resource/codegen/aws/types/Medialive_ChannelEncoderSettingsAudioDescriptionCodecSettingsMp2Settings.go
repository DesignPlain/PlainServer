package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsMp2Settings struct {
	//
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	//
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sample rate in Hz.
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`
}
