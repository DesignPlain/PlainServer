package types

type Medialive_InputDestination struct {
	// A unique name for the location the RTMP stream is being pushed to.
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`
}
