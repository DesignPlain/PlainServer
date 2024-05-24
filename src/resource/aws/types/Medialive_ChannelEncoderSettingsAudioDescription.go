package types

type Medialive_ChannelEncoderSettingsAudioDescription struct {
	// Advanced audio normalization settings. See Audio Normalization Settings for more details.
	AudioNormalizationSettings Medialive_ChannelEncoderSettingsAudioDescriptionAudioNormalizationSettings `json:"audioNormalizationSettings,omitempty" yaml:"audioNormalizationSettings,omitempty"`

	// The name of the audio selector used as the source for this AudioDescription.
	AudioSelectorName string `json:"audioSelectorName,omitempty" yaml:"audioSelectorName,omitempty"`

	//
	RemixSettings Medialive_ChannelEncoderSettingsAudioDescriptionRemixSettings `json:"remixSettings,omitempty" yaml:"remixSettings,omitempty"`

	// Stream name RTMP destinations (URLs of type rtmp://)
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`

	// The name of this audio description.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Applies only if audioTypeControl is useConfigured. The values for audioType are defined in ISO-IEC 13818-1.
	AudioType string `json:"audioType,omitempty" yaml:"audioType,omitempty"`

	// Determined how audio type is determined.
	AudioTypeControl string `json:"audioTypeControl,omitempty" yaml:"audioTypeControl,omitempty"`

	// Settings to configure one or more solutions that insert audio watermarks in the audio encode. See Audio Watermark Settings for more details.
	AudioWatermarkSettings Medialive_ChannelEncoderSettingsAudioDescriptionAudioWatermarkSettings `json:"audioWatermarkSettings,omitempty" yaml:"audioWatermarkSettings,omitempty"`

	// Audio codec settings. See Audio Codec Settings for more details.
	CodecSettings Medialive_ChannelEncoderSettingsAudioDescriptionCodecSettings `json:"codecSettings,omitempty" yaml:"codecSettings,omitempty"`

	// Selects a specific three-letter language code from within an audio source.
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	//
	LanguageCodeControl string `json:"languageCodeControl,omitempty" yaml:"languageCodeControl,omitempty"`
}
