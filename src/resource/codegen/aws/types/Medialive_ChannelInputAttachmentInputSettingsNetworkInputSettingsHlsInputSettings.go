package types

type Medialive_ChannelInputAttachmentInputSettingsNetworkInputSettingsHlsInputSettings struct {
	// The bitrate is specified in bits per second, as in an HLS manifest.
	Bandwidth int `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`

	// Buffer segments.
	BufferSegments int `json:"bufferSegments,omitempty" yaml:"bufferSegments,omitempty"`

	// The number of consecutive times that attempts to read a manifest or segment must fail before the input is considered unavailable.
	Retries int `json:"retries,omitempty" yaml:"retries,omitempty"`

	// The number of seconds between retries when an attempt to read a manifest or segment fails.
	RetryInterval int `json:"retryInterval,omitempty" yaml:"retryInterval,omitempty"`

	//
	Scte35Source string `json:"scte35Source,omitempty" yaml:"scte35Source,omitempty"`
}
