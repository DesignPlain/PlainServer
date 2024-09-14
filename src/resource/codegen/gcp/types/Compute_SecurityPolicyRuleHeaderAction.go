package types

type Compute_SecurityPolicyRuleHeaderAction struct {
	// The list of request headers to add or overwrite if they're already present. Structure is documented below.
	RequestHeadersToAdds []Compute_SecurityPolicyRuleHeaderActionRequestHeadersToAdd `json:"requestHeadersToAdds,omitempty" yaml:"requestHeadersToAdds,omitempty"`
}
