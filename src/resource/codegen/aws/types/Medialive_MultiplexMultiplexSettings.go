package types

type Medialive_MultiplexMultiplexSettings struct {
	// Maximum video buffer delay.
	MaximumVideoBufferDelayMilliseconds int `json:"maximumVideoBufferDelayMilliseconds,omitempty" yaml:"maximumVideoBufferDelayMilliseconds,omitempty"`

	// Transport stream bit rate.
	TransportStreamBitrate int `json:"transportStreamBitrate,omitempty" yaml:"transportStreamBitrate,omitempty"`

	// Unique ID for each multiplex.
	TransportStreamId int `json:"transportStreamId,omitempty" yaml:"transportStreamId,omitempty"`

	// Transport stream reserved bit rate.
	TransportStreamReservedBitrate int `json:"transportStreamReservedBitrate,omitempty" yaml:"transportStreamReservedBitrate,omitempty"`
}
