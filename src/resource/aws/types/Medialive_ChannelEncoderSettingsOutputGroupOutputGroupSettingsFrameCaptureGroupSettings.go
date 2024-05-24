package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettings struct {
	// A director and base filename where archive files should be written. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	//
	FrameCaptureCdnSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettings `json:"frameCaptureCdnSettings,omitempty" yaml:"frameCaptureCdnSettings,omitempty"`
}
