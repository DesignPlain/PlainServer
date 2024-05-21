package types

type Storage_getBucketCor struct {
	// The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" yaml:"maxAgeSeconds,omitempty"`

	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "-" is permitted in the list of methods, and means "any method".
	Methods []string `json:"methods,omitempty" yaml:"methods,omitempty"`

	// The list of Origins eligible to receive CORS response headers. Note: "-" is permitted in the list of origins, and means "any Origin".
	Origins []string `json:"origins,omitempty" yaml:"origins,omitempty"`

	// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
	ResponseHeaders []string `json:"responseHeaders,omitempty" yaml:"responseHeaders,omitempty"`
}
