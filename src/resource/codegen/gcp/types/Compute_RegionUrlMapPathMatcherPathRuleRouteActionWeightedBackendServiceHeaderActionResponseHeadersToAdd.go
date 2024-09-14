package types

type Compute_RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceHeaderActionResponseHeadersToAdd struct {
	// The name of the header.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// The value of the header to add.
	HeaderValue string `json:"headerValue,omitempty" yaml:"headerValue,omitempty"`

	/*
	   If false, headerValue is appended to any values that already exist for the header. If true, headerValue is set for the header, discarding any values that were set for that header.
	   The default value is false.
	*/
	Replace bool `json:"replace,omitempty" yaml:"replace,omitempty"`
}
