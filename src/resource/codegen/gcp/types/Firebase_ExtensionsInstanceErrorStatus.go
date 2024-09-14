package types

type Firebase_ExtensionsInstanceErrorStatus struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code int `json:"code,omitempty" yaml:"code,omitempty"`

	// A list of messages that carry the error details.
	Details []map[string]string `json:"details,omitempty" yaml:"details,omitempty"`

	// A developer-facing error message, which should be in English.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}
