package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsEac3Settings struct {
	//
	DcFilter string `json:"dcFilter,omitempty" yaml:"dcFilter,omitempty"`

	//
	LoRoCenterMixLevel float64 `json:"loRoCenterMixLevel,omitempty" yaml:"loRoCenterMixLevel,omitempty"`

	//
	LtRtSurroundMixLevel float64 `json:"ltRtSurroundMixLevel,omitempty" yaml:"ltRtSurroundMixLevel,omitempty"`

	//
	StereoDownmix string `json:"stereoDownmix,omitempty" yaml:"stereoDownmix,omitempty"`

	//
	SurroundMode string `json:"surroundMode,omitempty" yaml:"surroundMode,omitempty"`

	// Specifies the bitstream mode (bsmod) for the emitted AC-3 stream.
	BitstreamMode string `json:"bitstreamMode,omitempty" yaml:"bitstreamMode,omitempty"`

	//
	Dialnorm int `json:"dialnorm,omitempty" yaml:"dialnorm,omitempty"`

	// Sets the Dolby dynamic range compression profile.
	DrcLine string `json:"drcLine,omitempty" yaml:"drcLine,omitempty"`

	// When set to enabled, applies a 120Hz lowpass filter to the LFE channel prior to encoding.
	LfeFilter string `json:"lfeFilter,omitempty" yaml:"lfeFilter,omitempty"`

	// Metadata control.
	MetadataControl string `json:"metadataControl,omitempty" yaml:"metadataControl,omitempty"`

	//
	PassthroughControl string `json:"passthroughControl,omitempty" yaml:"passthroughControl,omitempty"`

	// Sets the attenuation control.
	AttenuationControl string `json:"attenuationControl,omitempty" yaml:"attenuationControl,omitempty"`

	// Average bitrate in bits/second.
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	//
	LfeControl string `json:"lfeControl,omitempty" yaml:"lfeControl,omitempty"`

	//
	LtRtCenterMixLevel float64 `json:"ltRtCenterMixLevel,omitempty" yaml:"ltRtCenterMixLevel,omitempty"`

	// Dolby Digital Plus coding mode.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Sets the profile for heavy Dolby dynamic range compression.
	DrcRf string `json:"drcRf,omitempty" yaml:"drcRf,omitempty"`

	//
	LoRoSurroundMixLevel float64 `json:"loRoSurroundMixLevel,omitempty" yaml:"loRoSurroundMixLevel,omitempty"`

	//
	PhaseControl string `json:"phaseControl,omitempty" yaml:"phaseControl,omitempty"`

	//
	SurroundExMode string `json:"surroundExMode,omitempty" yaml:"surroundExMode,omitempty"`
}
