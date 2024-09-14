package types

type Medialive_ChannelEncoderSettingsAudioDescription struct {
	// Determined how audio type is determined.
	AudioTypeControl string `json:"audioTypeControl,omitempty" yaml:"audioTypeControl,omitempty"`

	//
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	//
	LanguageCodeControl string `json:"languageCodeControl,omitempty" yaml:"languageCodeControl,omitempty"`

	// The name of this audio description.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	RemixSettings Medialive_ChannelEncoderSettingsAudioDescriptionRemixSettings `json:"remixSettings,omitempty" yaml:"remixSettings,omitempty"`

	// Applies only if audioTypeControl is useConfigured. The values for audioType are defined in ISO-IEC 13818-1.
	AudioType string `json:"audioType,omitempty" yaml:"audioType,omitempty"`

	// The name of the audio selector used as the source for this AudioDescription.
	AudioSelectorName string `json:"audioSelectorName,omitempty" yaml:"audioSelectorName,omitempty"`

	// Settings to configure one or more solutions that insert audio watermarks in the audio encode. See Audio Watermark Settings for more details.
	AudioWatermarkSettings Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettings `json:"audioWatermarkSettings,omitempty" yaml:"audioWatermarkSettings,omitempty"`

	// Audio codec settings. See Audio Codec Settings for more details.
	CodecSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettings `json:"codecSettings,omitempty" yaml:"codecSettings,omitempty"`

	// Stream name RTMP destinations (URLs of type rtmp://)
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`

	// Advanced audio normalization settings. See Audio Normalization Settings for more details.
	AudioNormalizationSettings Medialive_ChannelEncoderSettingsAudioDescriptionAudioNormalizationSettings `json:"audioNormalizationSettings,omitempty" yaml:"audioNormalizationSettings,omitempty"`
}
