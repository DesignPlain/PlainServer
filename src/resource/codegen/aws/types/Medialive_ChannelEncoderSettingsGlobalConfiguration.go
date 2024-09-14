package types

type Medialive_ChannelEncoderSettingsGlobalConfiguration struct {
	// Indicates whether the rate of frames emitted by the Live encoder should be paced by its system clock (which optionally may be locked to another source via NTP) or should be locked to the clock of the source that is providing the input stream.
	OutputTimingSource string `json:"outputTimingSource,omitempty" yaml:"outputTimingSource,omitempty"`

	// Adjusts video input buffer for streams with very low video framerates. This is commonly set to enabled for music channels with less than one video frame per second.
	SupportLowFramerateInputs string `json:"supportLowFramerateInputs,omitempty" yaml:"supportLowFramerateInputs,omitempty"`

	// Value to set the initial audio gain for the Live Event.
	InitialAudioGain int `json:"initialAudioGain,omitempty" yaml:"initialAudioGain,omitempty"`

	// Indicates the action to take when the current input completes (e.g. end-of-file). When switchAndLoopInputs is configured the encoder will restart at the beginning of the first input. When “none” is configured the encoder will transcode either black, a solid color, or a user specified slate images per the “Input Loss Behavior” configuration until the next input switch occurs (which is controlled through the Channel Schedule API).
	InputEndAction string `json:"inputEndAction,omitempty" yaml:"inputEndAction,omitempty"`

	// Settings for system actions when input is lost. See Input Loss Behavior for more details.
	InputLossBehavior Medialive_ChannelEncoderSettingsGlobalConfigurationInputLossBehavior `json:"inputLossBehavior,omitempty" yaml:"inputLossBehavior,omitempty"`

	// Indicates how MediaLive pipelines are synchronized. PIPELINE\_LOCKING - MediaLive will attempt to synchronize the output of each pipeline to the other. EPOCH\_LOCKING - MediaLive will attempt to synchronize the output of each pipeline to the Unix epoch.
	OutputLockingMode string `json:"outputLockingMode,omitempty" yaml:"outputLockingMode,omitempty"`
}
