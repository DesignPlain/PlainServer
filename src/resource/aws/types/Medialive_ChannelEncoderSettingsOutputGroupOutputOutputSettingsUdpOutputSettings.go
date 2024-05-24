package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettings struct {
	//
	FecOutputSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsFecOutputSettings `json:"fecOutputSettings,omitempty" yaml:"fecOutputSettings,omitempty"`

	// UDP output buffering in milliseconds.
	BufferMsec int `json:"bufferMsec,omitempty" yaml:"bufferMsec,omitempty"`

	// UDP container settings. See Container Settings for more details.
	ContainerSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsContainerSettings `json:"containerSettings,omitempty" yaml:"containerSettings,omitempty"`

	// Destination address and port number for RTP or UDP packets. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`
}
