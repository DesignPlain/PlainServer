package types

type Cloudfront_RealtimeLogConfigEndpoint struct {
	// The Amazon Kinesis data stream configuration.
	KinesisStreamConfig Cloudfront_RealtimeLogConfigEndpointKinesisStreamConfig `json:"kinesisStreamConfig,omitempty" yaml:"kinesisStreamConfig,omitempty"`

	// The type of data stream where real-time log data is sent. The only valid value is `Kinesis`.
	StreamType string `json:"streamType,omitempty" yaml:"streamType,omitempty"`
}
