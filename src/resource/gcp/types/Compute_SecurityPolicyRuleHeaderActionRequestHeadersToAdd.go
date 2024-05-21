package types

type Compute_SecurityPolicyRuleHeaderActionRequestHeadersToAdd struct {
	// The name of the header to set.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// The value to set the named header to.
	HeaderValue string `json:"headerValue,omitempty" yaml:"headerValue,omitempty"`
}
