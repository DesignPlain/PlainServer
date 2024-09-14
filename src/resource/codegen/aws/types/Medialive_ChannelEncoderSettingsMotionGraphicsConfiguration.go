package types

type Medialive_ChannelEncoderSettingsMotionGraphicsConfiguration struct {
	// Motion Graphics Insertion.
	MotionGraphicsInsertion string `json:"motionGraphicsInsertion,omitempty" yaml:"motionGraphicsInsertion,omitempty"`

	// Motion Graphics Settings. See Motion Graphics Settings for more details.
	MotionGraphicsSettings Medialive_ChannelEncoderSettingsMotionGraphicsConfigurationMotionGraphicsSettings `json:"motionGraphicsSettings,omitempty" yaml:"motionGraphicsSettings,omitempty"`
}
