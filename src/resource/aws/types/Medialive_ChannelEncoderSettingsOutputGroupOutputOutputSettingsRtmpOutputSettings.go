package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettings struct {
	// Setting to allow self signed or verified RTMP certificates.
	CertificateMode string `json:"certificateMode,omitempty" yaml:"certificateMode,omitempty"`

	// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
	ConnectionRetryInterval int `json:"connectionRetryInterval,omitempty" yaml:"connectionRetryInterval,omitempty"`

	// The RTMP endpoint excluding the stream name. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsRtmpOutputSettingsDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Number of retry attempts.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`
}
