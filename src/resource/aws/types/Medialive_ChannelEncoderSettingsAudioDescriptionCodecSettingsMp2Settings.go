package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsMp2Settings struct {
	// Mono, Stereo, or 5.1 channel layout.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sample rate in Hz.
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`

	// Average bitrate in bits/second.
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`
}
