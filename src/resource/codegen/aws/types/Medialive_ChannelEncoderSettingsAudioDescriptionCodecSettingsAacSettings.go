package types

type Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettingsAacSettings struct {
	// The rate control mode.
	RateControlMode string `json:"rateControlMode,omitempty" yaml:"rateControlMode,omitempty"`

	// Sets LATM/LOAS AAC output for raw containers.
	RawFormat string `json:"rawFormat,omitempty" yaml:"rawFormat,omitempty"`

	// Use MPEG-2 AAC audio instead of MPEG-4 AAC audio for raw or MPEG-2 Transport Stream containers.
	Spec string `json:"spec,omitempty" yaml:"spec,omitempty"`

	// VBR Quality Level - Only used if rateControlMode is VBR.
	VbrQuality string `json:"vbrQuality,omitempty" yaml:"vbrQuality,omitempty"`

	// Mono, Stereo, or 5.1 channel layout.
	CodingMode string `json:"codingMode,omitempty" yaml:"codingMode,omitempty"`

	// Set to "broadcasterMixedAd" when input contains pre-mixed main audio + AD (narration) as a stereo pair.
	InputType string `json:"inputType,omitempty" yaml:"inputType,omitempty"`

	// Sample rate in Hz.
	SampleRate float64 `json:"sampleRate,omitempty" yaml:"sampleRate,omitempty"`

	// Average bitrate in bits/second.
	Bitrate float64 `json:"bitrate,omitempty" yaml:"bitrate,omitempty"`

	// AAC profile.
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`
}
