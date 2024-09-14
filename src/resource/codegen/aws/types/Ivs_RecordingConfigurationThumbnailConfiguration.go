package types

type Ivs_RecordingConfigurationThumbnailConfiguration struct {
	// Thumbnail recording mode. Valid values: `DISABLED`, `INTERVAL`.
	RecordingMode string `json:"recordingMode,omitempty" yaml:"recordingMode,omitempty"`

	// The targeted thumbnail-generation interval in seconds.
	TargetIntervalSeconds int `json:"targetIntervalSeconds,omitempty" yaml:"targetIntervalSeconds,omitempty"`
}
