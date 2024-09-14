package types

type Medialive_ChannelEncoderSettingsAvailBlanking struct {
	// Blanking image to be used. See Avail Blanking Image for more details.
	AvailBlankingImage Medialive_ChannelEncoderSettingsAvailBlankingAvailBlankingImage `json:"availBlankingImage,omitempty" yaml:"availBlankingImage,omitempty"`

	// When set to enabled, causes video, audio and captions to be blanked when insertion metadata is added.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
