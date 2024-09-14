package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelector struct {
	//
	SelectorSettings Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettings `json:"selectorSettings,omitempty" yaml:"selectorSettings,omitempty"`

	//
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   Name of the Channel.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
