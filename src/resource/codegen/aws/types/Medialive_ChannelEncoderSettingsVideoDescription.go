package types

type Medialive_ChannelEncoderSettingsVideoDescription struct {
	// Behavior on how to scale.
	ScalingBehavior string `json:"scalingBehavior,omitempty" yaml:"scalingBehavior,omitempty"`

	// Changes the strength of the anti-alias filter used for scaling.
	Sharpness int `json:"sharpness,omitempty" yaml:"sharpness,omitempty"`

	// Output video width in pixels.
	Width int `json:"width,omitempty" yaml:"width,omitempty"`

	// The video codec settings. See Video Codec Settings for more details.
	CodecSettings Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettings `json:"codecSettings,omitempty" yaml:"codecSettings,omitempty"`

	// Output video height in pixels.
	Height int `json:"height,omitempty" yaml:"height,omitempty"`

	// The name of the video description.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Indicate how to respond to the AFD values that might be in the input video.
	RespondToAfd string `json:"respondToAfd,omitempty" yaml:"respondToAfd,omitempty"`
}
