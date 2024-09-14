package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettings struct {
	//
	FragmentLength int `json:"fragmentLength,omitempty" yaml:"fragmentLength,omitempty"`

	// Number of retry attempts.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`

	// Number of seconds to wait until a restart is initiated.
	RestartDelay int `json:"restartDelay,omitempty" yaml:"restartDelay,omitempty"`

	//
	TimestampOffset string `json:"timestampOffset,omitempty" yaml:"timestampOffset,omitempty"`

	// Setting to allow self signed or verified RTMP certificates.
	CertificateMode string `json:"certificateMode,omitempty" yaml:"certificateMode,omitempty"`

	// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
	ConnectionRetryInterval int `json:"connectionRetryInterval,omitempty" yaml:"connectionRetryInterval,omitempty"`

	//
	EventId string `json:"eventId,omitempty" yaml:"eventId,omitempty"`

	//
	SendDelayMs int `json:"sendDelayMs,omitempty" yaml:"sendDelayMs,omitempty"`

	//
	AcquisitionPointId string `json:"acquisitionPointId,omitempty" yaml:"acquisitionPointId,omitempty"`

	//
	AudioOnlyTimecodeControl string `json:"audioOnlyTimecodeControl,omitempty" yaml:"audioOnlyTimecodeControl,omitempty"`

	//
	EventStopBehavior string `json:"eventStopBehavior,omitempty" yaml:"eventStopBehavior,omitempty"`

	//
	SegmentationMode string `json:"segmentationMode,omitempty" yaml:"segmentationMode,omitempty"`

	//
	TimestampOffsetMode string `json:"timestampOffsetMode,omitempty" yaml:"timestampOffsetMode,omitempty"`

	//
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	//
	SparseTrackType string `json:"sparseTrackType,omitempty" yaml:"sparseTrackType,omitempty"`

	//
	StreamManifestBehavior string `json:"streamManifestBehavior,omitempty" yaml:"streamManifestBehavior,omitempty"`

	//
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	EventIdMode string `json:"eventIdMode,omitempty" yaml:"eventIdMode,omitempty"`

	//
	FilecacheDuration int `json:"filecacheDuration,omitempty" yaml:"filecacheDuration,omitempty"`
}
