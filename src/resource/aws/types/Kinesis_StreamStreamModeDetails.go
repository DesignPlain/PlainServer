package types

type Kinesis_StreamStreamModeDetails struct {
	// Specifies the capacity mode of the stream. Must be either `PROVISIONED` or `ON_DEMAND`.
	StreamMode string `json:"streamMode,omitempty" yaml:"streamMode,omitempty"`
}
