package types

type Cloudfront_ResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity struct {
	// A number that CloudFront uses as the value for the `max-age` directive in the `Strict-Transport-Security` HTTP response header.
	AccessControlMaxAgeSec int `json:"accessControlMaxAgeSec,omitempty" yaml:"accessControlMaxAgeSec,omitempty"`

	// Whether CloudFront includes the `includeSubDomains` directive in the `Strict-Transport-Security` HTTP response header.
	IncludeSubdomains bool `json:"includeSubdomains,omitempty" yaml:"includeSubdomains,omitempty"`

	// Whether CloudFront overrides the `Strict-Transport-Security` HTTP response header received from the origin with the one specified in this response headers policy.
	Override bool `json:"override,omitempty" yaml:"override,omitempty"`

	// Whether CloudFront includes the `preload` directive in the `Strict-Transport-Security` HTTP response header.
	Preload bool `json:"preload,omitempty" yaml:"preload,omitempty"`
}
