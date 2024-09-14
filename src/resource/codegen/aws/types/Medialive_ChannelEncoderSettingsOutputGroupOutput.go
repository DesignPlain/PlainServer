package types

type Medialive_ChannelEncoderSettingsOutputGroupOutput struct {
	// The names of the audio descriptions used as audio sources for the output.
	AudioDescriptionNames []string `json:"audioDescriptionNames,omitempty" yaml:"audioDescriptionNames,omitempty"`

	// The names of the caption descriptions used as caption sources for the output.
	CaptionDescriptionNames []string `json:"captionDescriptionNames,omitempty" yaml:"captionDescriptionNames,omitempty"`

	// The name used to identify an output.
	OutputName string `json:"outputName,omitempty" yaml:"outputName,omitempty"`

	// Settings for output. See Output Settings for more details.
	OutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettings `json:"outputSettings,omitempty" yaml:"outputSettings,omitempty"`

	// The name of the video description used as video source for the output.
	VideoDescriptionName string `json:"videoDescriptionName,omitempty" yaml:"videoDescriptionName,omitempty"`
}
