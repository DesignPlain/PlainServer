package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsAc3Settings struct {
	// Average bitrate in bits/second.
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	BitstreamMode string `json:"bitstreamMode,omitempty" yaml:"bitstreamMode,omitempty"`

	// Mono, Stereo, or 5.1 channel layout.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sets the dialnorm of the output.
	Dialnorm int `json:"dialnorm,omitempty" yaml:"dialnorm,omitempty"`

	// If set to filmStandard, adds dynamic range compression signaling to the output bitstream as defined in the Dolby Digital specification.
	DrcProfile string `json:"drcProfile,omitempty" yaml:"drcProfile,omitempty"`

	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	LfeFilter string `json:"lfeFilter,omitempty" yaml:"lfeFilter,omitempty"`

	// Metadata control.
	MetadataControl string `json:"metadataControl,omitempty" yaml:"metadataControl,omitempty"`
}
