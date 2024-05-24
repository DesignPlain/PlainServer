package types

type Medialive_ChannelInputAttachmentInputSettings struct {
	//
	CaptionSelectors []Medialive_ChannelInputAttachmentInputSettingsCaptionSelector `json:"captionSelectors,omitempty" yaml:"captionSelectors,omitempty"`

	// Adjusts the magnitude of filtering from 1 (minimal) to 5 (strongest).
	FilterStrength int `json:"filterStrength,omitempty" yaml:"filterStrength,omitempty"`

	// Input settings. See Network Input Settings for more details.
	NetworkInputSettings Medialive_ChannelInputAttachmentInputSettingsNetworkInputSettings `json:"networkInputSettings,omitempty" yaml:"networkInputSettings,omitempty"`

	// PID from which to read SCTE-35 messages.
	Scte35Pid int `json:"scte35Pid,omitempty" yaml:"scte35Pid,omitempty"`

	// Loop input if it is a file.
	SourceEndBehavior string `json:"sourceEndBehavior,omitempty" yaml:"sourceEndBehavior,omitempty"`

	//
	VideoSelector Medialive_ChannelInputAttachmentInputSettingsVideoSelector `json:"videoSelector,omitempty" yaml:"videoSelector,omitempty"`

	//
	AudioSelectors []Medialive_ChannelInputAttachmentInputSettingsAudioSelector `json:"audioSelectors,omitempty" yaml:"audioSelectors,omitempty"`

	// Enable or disable the deblock filter when filtering.
	DeblockFilter string `json:"deblockFilter,omitempty" yaml:"deblockFilter,omitempty"`

	// Enable or disable the denoise filter when filtering.
	DenoiseFilter string `json:"denoiseFilter,omitempty" yaml:"denoiseFilter,omitempty"`

	// Turns on the filter for the input.
	InputFilter string `json:"inputFilter,omitempty" yaml:"inputFilter,omitempty"`

	// Specifies whether to extract applicable ancillary data from a SMPTE-2038 source in the input.
	Smpte2038DataPreference string `json:"smpte2038DataPreference,omitempty" yaml:"smpte2038DataPreference,omitempty"`
}
