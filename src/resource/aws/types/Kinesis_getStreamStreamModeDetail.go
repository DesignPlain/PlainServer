package types

type Kinesis_getStreamStreamModeDetail struct {
	// Capacity mode of the stream. Either `ON_DEMAND` or `PROVISIONED`.
	StreamMode string `json:"streamMode,omitempty" yaml:"streamMode,omitempty"`
}
