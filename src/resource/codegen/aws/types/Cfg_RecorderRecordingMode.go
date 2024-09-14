package types

type Cfg_RecorderRecordingMode struct {
	// Default reecording frequency. `CONTINUOUS` or `DAILY`.
	RecordingFrequency string `json:"recordingFrequency,omitempty" yaml:"recordingFrequency,omitempty"`

	// Recording mode overrides. Detailed below.
	RecordingModeOverride Cfg_RecorderRecordingModeRecordingModeOverride `json:"recordingModeOverride,omitempty" yaml:"recordingModeOverride,omitempty"`
}
