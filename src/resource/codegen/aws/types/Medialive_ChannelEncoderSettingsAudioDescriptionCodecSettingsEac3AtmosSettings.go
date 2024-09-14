package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsEac3AtmosSettings struct {
	// Average bitrate in bits/second.
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	// Dolby Digital Plus with Dolby Atmos coding mode.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sets the dialnorm for the output.
	Dialnorm float64 `json:"dialnorm,omitempty" yaml:"dialnorm,omitempty"`

	// Sets the Dolby dynamic range compression profile.
	DrcLine string `json:"drcLine,omitempty" yaml:"drcLine,omitempty"`

	// Sets the profile for heavy Dolby dynamic range compression.
	DrcRf string `json:"drcRf,omitempty" yaml:"drcRf,omitempty"`

	// Height dimensional trim.
	HeightTrim float64 `json:"heightTrim,omitempty" yaml:"heightTrim,omitempty"`

	// Surround dimensional trim.
	SurroundTrim float64 `json:"surroundTrim,omitempty" yaml:"surroundTrim,omitempty"`
}
