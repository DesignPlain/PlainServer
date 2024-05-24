package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettings struct {
	// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
	ConnectionRetryInterval int `json:"connectionRetryInterval,omitempty" yaml:"connectionRetryInterval,omitempty"`

	// A director and base filename where archive files should be written. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsMsSmoothGroupSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Number of seconds to wait until a restart is initiated.
	RestartDelay int `json:"restartDelay,omitempty" yaml:"restartDelay,omitempty"`

	//
	StreamManifestBehavior string `json:"streamManifestBehavior,omitempty" yaml:"streamManifestBehavior,omitempty"`

	//
	AcquisitionPointId string `json:"acquisitionPointId,omitempty" yaml:"acquisitionPointId,omitempty"`

	//
	SendDelayMs int `json:"sendDelayMs,omitempty" yaml:"sendDelayMs,omitempty"`

	//
	SparseTrackType string `json:"sparseTrackType,omitempty" yaml:"sparseTrackType,omitempty"`

	//
	EventId string `json:"eventId,omitempty" yaml:"eventId,omitempty"`

	//
	FilecacheDuration int `json:"filecacheDuration,omitempty" yaml:"filecacheDuration,omitempty"`

	//
	SegmentationMode string `json:"segmentationMode,omitempty" yaml:"segmentationMode,omitempty"`

	//
	TimestampOffsetMode string `json:"timestampOffsetMode,omitempty" yaml:"timestampOffsetMode,omitempty"`

	// Controls the behavior of the RTMP group if input becomes unavailable.
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	// Number of retry attempts.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`

	//
	TimestampOffset string `json:"timestampOffset,omitempty" yaml:"timestampOffset,omitempty"`

	//
	AudioOnlyTimecodeControl string `json:"audioOnlyTimecodeControl,omitempty" yaml:"audioOnlyTimecodeControl,omitempty"`

	// Setting to allow self signed or verified RTMP certificates.
	CertificateMode string `json:"certificateMode,omitempty" yaml:"certificateMode,omitempty"`

	//
	EventIdMode string `json:"eventIdMode,omitempty" yaml:"eventIdMode,omitempty"`

	//
	EventStopBehavior string `json:"eventStopBehavior,omitempty" yaml:"eventStopBehavior,omitempty"`

	//
	FragmentLength int `json:"fragmentLength,omitempty" yaml:"fragmentLength,omitempty"`
}
