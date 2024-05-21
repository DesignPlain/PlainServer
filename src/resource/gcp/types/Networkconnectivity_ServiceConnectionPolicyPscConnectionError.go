package types

type Networkconnectivity_ServiceConnectionPolicyPscConnectionError struct {
	/*
	   (Output)
	   A list of messages that carry the error details.
	*/
	Details []map[string]string `json:"details,omitempty" yaml:"details,omitempty"`

	// A developer-facing error message.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The status code, which should be an enum value of [google.rpc.Code][].
	Code int `json:"code,omitempty" yaml:"code,omitempty"`
}
