package types

type Medialive_ChannelEncoderSettingsVideoDescriptionCodecSettingsFrameCaptureSettings struct {
	// The frequency at which to capture frames for inclusion in the output.
	CaptureInterval int `json:"captureInterval,omitempty" yaml:"captureInterval,omitempty"`

	// Unit for the frame capture interval.
	CaptureIntervalUnits string `json:"captureIntervalUnits,omitempty" yaml:"captureIntervalUnits,omitempty"`
}
