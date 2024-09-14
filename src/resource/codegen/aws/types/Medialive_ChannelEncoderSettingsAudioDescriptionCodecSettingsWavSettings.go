package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsWavSettings struct {
	//
	BitDepth float64 `json:"bitDepth,omitempty" yaml:"bitDepth,omitempty"`

	//
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sample rate in Hz.
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`
}
