package types

type Networksecurity_AuthorizationPolicyRuleDestination struct {
	// List of host names to match. Matched against the ":authority" header in http requests. At least one host should match. Each host can be an exact match, or a prefix match (example "mydomain.-") or a suffix match (example "-.myorg.com") or a presence (any) match "-".
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	/*
	   Match against key:value pair in http header. Provides a flexible match based on HTTP headers, for potentially advanced use cases. At least one header should match.
	   Avoid using header matches to make authorization decisions unless there is a strong guarantee that requests arrive through a trusted client or proxy.
	   Structure is documented below.
	*/
	HttpHeaderMatch Networksecurity_AuthorizationPolicyRuleDestinationHttpHeaderMatch `json:"httpHeaderMatch,omitempty" yaml:"httpHeaderMatch,omitempty"`

	// A list of HTTP methods to match. At least one method should match. Should not be set for gRPC services.
	Methods []string `json:"methods,omitempty" yaml:"methods,omitempty"`

	// List of destination ports to match. At least one port should match.
	Ports []int `json:"ports,omitempty" yaml:"ports,omitempty"`
}
