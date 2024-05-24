package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type ResponseHeadersPolicy struct {
	// The current version of the response headers policy.
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	// A unique name to identify the response headers policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A configuration for a set of HTTP headers to remove from the HTTP response. Object that contains an attribute `items` that contains a list of headers. See Remove Header for more information.
	RemoveHeadersConfig types.Cloudfront_ResponseHeadersPolicyRemoveHeadersConfig `json:"removeHeadersConfig,omitempty" yaml:"removeHeadersConfig,omitempty"`

	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig types.Cloudfront_ResponseHeadersPolicySecurityHeadersConfig `json:"securityHeadersConfig,omitempty" yaml:"securityHeadersConfig,omitempty"`

	// A configuration for enabling the Server-Timing header in HTTP responses sent from CloudFront. See Server Timing Headers Config for more information.
	ServerTimingHeadersConfig types.Cloudfront_ResponseHeadersPolicyServerTimingHeadersConfig `json:"serverTimingHeadersConfig,omitempty" yaml:"serverTimingHeadersConfig,omitempty"`

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig types.Cloudfront_ResponseHeadersPolicyCorsConfig `json:"corsConfig,omitempty" yaml:"corsConfig,omitempty"`

	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig types.Cloudfront_ResponseHeadersPolicyCustomHeadersConfig `json:"customHeadersConfig,omitempty" yaml:"customHeadersConfig,omitempty"`
}
