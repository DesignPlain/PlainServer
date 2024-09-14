package types

type Firebase_HostingCustomDomainIssue struct {
	// The status code, which should be an enum value of `google.rpc.Code`
	Code int `json:"code,omitempty" yaml:"code,omitempty"`

	// A list of messages that carry the error details.
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	// Error message
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}
