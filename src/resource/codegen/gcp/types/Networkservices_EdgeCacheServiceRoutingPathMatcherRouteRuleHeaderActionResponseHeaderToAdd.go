package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAdd struct {
	// Whether to replace all existing headers with the same name.
	Replace bool `json:"replace,omitempty" yaml:"replace,omitempty"`

	// The name of the header to add.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// The value of the header to add.
	HeaderValue string `json:"headerValue,omitempty" yaml:"headerValue,omitempty"`
}
