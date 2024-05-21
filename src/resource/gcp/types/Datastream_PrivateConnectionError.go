package types

type Datastream_PrivateConnectionError struct {
	// A list of messages that carry the error details.
	Details map[string]string `json:"details,omitempty" yaml:"details,omitempty"`

	// A message containing more information about the error that occurred.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}
