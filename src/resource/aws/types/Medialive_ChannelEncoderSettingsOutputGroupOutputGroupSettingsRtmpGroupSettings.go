package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsRtmpGroupSettings struct {
	// Cache length in seconds, is used to calculate buffer size.
	CacheLength int `json:"cacheLength,omitempty" yaml:"cacheLength,omitempty"`

	// Controls the types of data that passes to onCaptionInfo outputs.
	CaptionData string `json:"captionData,omitempty" yaml:"captionData,omitempty"`

	// Controls the behavior of the RTMP group if input becomes unavailable.
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	// Number of seconds to wait until a restart is initiated.
	RestartDelay int `json:"restartDelay,omitempty" yaml:"restartDelay,omitempty"`

	// The ad marker type for this output group.
	AdMarkers []string `json:"adMarkers,omitempty" yaml:"adMarkers,omitempty"`

	// Authentication scheme to use when connecting with CDN.
	AuthenticationScheme string `json:"authenticationScheme,omitempty" yaml:"authenticationScheme,omitempty"`

	// Controls behavior when content cache fills up.
	CacheFullBehavior string `json:"cacheFullBehavior,omitempty" yaml:"cacheFullBehavior,omitempty"`
}
