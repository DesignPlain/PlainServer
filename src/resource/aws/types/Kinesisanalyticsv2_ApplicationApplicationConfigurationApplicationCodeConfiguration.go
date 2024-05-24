package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfiguration struct {
	// The location and type of the application code.
	CodeContent Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent `json:"codeContent,omitempty" yaml:"codeContent,omitempty"`

	// Specifies whether the code content is in text or zip format. Valid values: `PLAINTEXT`, `ZIPFILE`.
	CodeContentType string `json:"codeContentType,omitempty" yaml:"codeContentType,omitempty"`
}
