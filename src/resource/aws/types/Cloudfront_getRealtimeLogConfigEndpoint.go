package types

type Cloudfront_getRealtimeLogConfigEndpoint struct {
	// (Required) Amazon Kinesis data stream configuration.
	KinesisStreamConfigs []Cloudfront_getRealtimeLogConfigEndpointKinesisStreamConfig `json:"kinesisStreamConfigs,omitempty" yaml:"kinesisStreamConfigs,omitempty"`

	// (Required) Type of data stream where real-time log data is sent. The only valid value is `Kinesis`.
	StreamType string `json:"streamType,omitempty" yaml:"streamType,omitempty"`
}
