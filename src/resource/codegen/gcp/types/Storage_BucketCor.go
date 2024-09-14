package types

type Storage_BucketCor struct {
	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "-" is permitted in the list of methods, and means "any method".
	Methods []string `json:"methods,omitempty" yaml:"methods,omitempty"`

	// The list of [Origins](https://tools.ietf.org/html/rfc6454) eligible to receive CORS response headers. Note: "-" is permitted in the list of origins, and means "any Origin".
	Origins []string `json:"origins,omitempty" yaml:"origins,omitempty"`

	// The list of HTTP headers other than the [simple response headers](https://www.w3.org/TR/cors/#simple-response-header) to give permission for the user-agent to share across domains.
	ResponseHeaders []string `json:"responseHeaders,omitempty" yaml:"responseHeaders,omitempty"`

	// The value, in seconds, to return in the [Access-Control-Max-Age header](https://www.w3.org/TR/cors/#access-control-max-age-response-header) used in preflight responses.
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" yaml:"maxAgeSeconds,omitempty"`
}
