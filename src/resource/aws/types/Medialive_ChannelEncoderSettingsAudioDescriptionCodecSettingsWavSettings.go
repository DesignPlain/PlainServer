package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettings struct {
	// Sample rate in Hz.
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`

	//
	BitDepth float64 `json:"bitDepth,omitempty" yaml:"bitDepth,omitempty"`

	// Mono, Stereo, or 5.1 channel layout.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`
}
