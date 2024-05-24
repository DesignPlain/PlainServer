package types

type Medialive_ChannelEncoderSettingsAudioDescriptionRemixSettings struct {
	//
	ChannelMappings []Medialive_ChannelEncoderSettingsAudioDescriptionRemixSettingsChannelMapping `json:"channelMappings,omitempty" yaml:"channelMappings,omitempty"`

	//
	ChannelsIn int `json:"channelsIn,omitempty" yaml:"channelsIn,omitempty"`

	//
	ChannelsOut int `json:"channelsOut,omitempty" yaml:"channelsOut,omitempty"`
}
