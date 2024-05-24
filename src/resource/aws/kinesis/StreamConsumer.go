package kinesis

type StreamConsumer struct {
	// Name of the stream consumer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amazon Resource Name (ARN) of the data stream the consumer is registered with.
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`
}
