package types

type Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettings struct {
	// Audio HLS Rendition Selection. See Audio HLS Rendition Selection for more details.
	AudioHlsRenditionSelection Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioHlsRenditionSelection `json:"audioHlsRenditionSelection,omitempty" yaml:"audioHlsRenditionSelection,omitempty"`

	// Audio Language Selection. See Audio Language Selection for more details.
	AudioLanguageSelection Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioLanguageSelection `json:"audioLanguageSelection,omitempty" yaml:"audioLanguageSelection,omitempty"`

	// Audio Pid Selection. See Audio PID Selection for more details.
	AudioPidSelection Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioPidSelection `json:"audioPidSelection,omitempty" yaml:"audioPidSelection,omitempty"`

	// Audio Track Selection. See Audio Track Selection for more details.
	AudioTrackSelection Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioTrackSelection `json:"audioTrackSelection,omitempty" yaml:"audioTrackSelection,omitempty"`
}
