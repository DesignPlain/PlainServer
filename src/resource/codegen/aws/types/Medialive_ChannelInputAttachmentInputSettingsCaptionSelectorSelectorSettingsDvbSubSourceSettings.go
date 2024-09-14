package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsDvbSubSourceSettings struct {
	// When using DVB-Sub with Burn-In or SMPTE-TT, use this PID for the source content. Unused for DVB-Sub passthrough. All DVB-Sub content is passed through, regardless of selectors.
	Pid int `json:"pid,omitempty" yaml:"pid,omitempty"`

	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage string `json:"ocrLanguage,omitempty" yaml:"ocrLanguage,omitempty"`
}
