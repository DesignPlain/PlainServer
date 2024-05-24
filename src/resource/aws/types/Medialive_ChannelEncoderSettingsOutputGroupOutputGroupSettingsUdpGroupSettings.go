package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsUdpGroupSettings struct {
	// Specifies behavior of last resort when input video os lost.
	InputLossAction string `json:"inputLossAction,omitempty" yaml:"inputLossAction,omitempty"`

	// Indicates ID3 frame that has the timecode.
	TimedMetadataId3Frame string `json:"timedMetadataId3Frame,omitempty" yaml:"timedMetadataId3Frame,omitempty"`

	//
	TimedMetadataId3Period int `json:"timedMetadataId3Period,omitempty" yaml:"timedMetadataId3Period,omitempty"`
}
