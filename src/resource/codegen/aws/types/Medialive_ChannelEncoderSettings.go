package types

type Medialive_ChannelEncoderSettings struct {
	// Nielsen configuration settings. See Nielsen Configuration for more details.
	NielsenConfiguration Medialive_ChannelEncoderSettingsNielsenConfiguration `json:"nielsenConfiguration,omitempty" yaml:"nielsenConfiguration,omitempty"`

	// Contains settings used to acquire and adjust timecode information from inputs. See Timecode Config for more details.
	TimecodeConfig Medialive_ChannelEncoderSettingsTimecodeConfig `json:"timecodeConfig,omitempty" yaml:"timecodeConfig,omitempty"`

	// Audio descriptions for the channel. See Audio Descriptions for more details.
	AudioDescriptions []Medialive_ChannelEncoderSettingsAudioDescription `json:"audioDescriptions,omitempty" yaml:"audioDescriptions,omitempty"`

	// Settings for ad avail blanking. See Avail Blanking for more details.
	AvailBlanking Medialive_ChannelEncoderSettingsAvailBlanking `json:"availBlanking,omitempty" yaml:"availBlanking,omitempty"`

	// Configuration settings that apply to the event as a whole. See Global Configuration for more details.
	GlobalConfiguration Medialive_ChannelEncoderSettingsGlobalConfiguration `json:"globalConfiguration,omitempty" yaml:"globalConfiguration,omitempty"`

	// Settings for motion graphics. See Motion Graphics Configuration for more details.
	MotionGraphicsConfiguration Medialive_ChannelEncoderSettingsMotionGraphicsConfiguration `json:"motionGraphicsConfiguration,omitempty" yaml:"motionGraphicsConfiguration,omitempty"`

	// Caption Descriptions. See Caption Descriptions for more details.
	CaptionDescriptions []Medialive_ChannelEncoderSettingsCaptionDescription `json:"captionDescriptions,omitempty" yaml:"captionDescriptions,omitempty"`

	// Output groups for the channel. See Output Groups for more details.
	OutputGroups []Medialive_ChannelEncoderSettingsOutputGroup `json:"outputGroups,omitempty" yaml:"outputGroups,omitempty"`

	// Video Descriptions. See Video Descriptions for more details.
	VideoDescriptions []Medialive_ChannelEncoderSettingsVideoDescription `json:"videoDescriptions,omitempty" yaml:"videoDescriptions,omitempty"`
}
