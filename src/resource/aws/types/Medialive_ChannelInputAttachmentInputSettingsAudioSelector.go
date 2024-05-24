package types

type Medialive_ChannelInputAttachmentInputSettingsAudioSelector struct {
	/*
	   Name of the Channel.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The audio selector settings. See Audio Selector Settings for more details.
	SelectorSettings Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettings `json:"selectorSettings,omitempty" yaml:"selectorSettings,omitempty"`
}
