package types

type Medialive_ChannelEncoderSettingsAudioDescriptionAudioNormalizationSettings struct {
	// Audio normalization algorithm to use. itu17701 conforms to the CALM Act specification, itu17702 to the EBU R-128 specification.
	Algorithm string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`

	// Algorithm control for the audio description.
	AlgorithmControl string `json:"algorithmControl,omitempty" yaml:"algorithmControl,omitempty"`

	// Target LKFS (loudness) to adjust volume to.
	TargetLkfs float64 `json:"targetLkfs,omitempty" yaml:"targetLkfs,omitempty"`
}
