package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsScte27SourceSettings struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage string `json:"ocrLanguage,omitempty" yaml:"ocrLanguage,omitempty"`

	// Selects a specific PID from within a source.
	Pid int `json:"pid,omitempty" yaml:"pid,omitempty"`
}
