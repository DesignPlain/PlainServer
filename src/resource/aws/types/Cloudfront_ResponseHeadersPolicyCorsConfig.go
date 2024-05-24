package types

type Cloudfront_ResponseHeadersPolicyCorsConfig struct {
	// A number that CloudFront uses as the value for the `Access-Control-Max-Age` HTTP response header.
	AccessControlMaxAgeSec int `json:"accessControlMaxAgeSec,omitempty" yaml:"accessControlMaxAgeSec,omitempty"`

	// A Boolean value that determines how CloudFront behaves for the HTTP response header.
	OriginOverride bool `json:"originOverride,omitempty" yaml:"originOverride,omitempty"`

	// A Boolean value that CloudFront uses as the value for the `Access-Control-Allow-Credentials` HTTP response header.
	AccessControlAllowCredentials bool `json:"accessControlAllowCredentials,omitempty" yaml:"accessControlAllowCredentials,omitempty"`

	// Object that contains an attribute `items` that contains a list of HTTP header names that CloudFront includes as values for the `Access-Control-Allow-Headers` HTTP response header.
	AccessControlAllowHeaders Cloudfront_ResponseHeadersPolicyCorsConfigAccessControlAllowHeaders `json:"accessControlAllowHeaders,omitempty" yaml:"accessControlAllowHeaders,omitempty"`

	// Object that contains an attribute `items` that contains a list of HTTP methods that CloudFront includes as values for the `Access-Control-Allow-Methods` HTTP response header. Valid values: `GET` | `POST` | `OPTIONS` | `PUT` | `DELETE` | `HEAD` | `ALL`
	AccessControlAllowMethods Cloudfront_ResponseHeadersPolicyCorsConfigAccessControlAllowMethods `json:"accessControlAllowMethods,omitempty" yaml:"accessControlAllowMethods,omitempty"`

	// Object that contains an attribute `items` that contains a list of origins that CloudFront can use as the value for the `Access-Control-Allow-Origin` HTTP response header.
	AccessControlAllowOrigins Cloudfront_ResponseHeadersPolicyCorsConfigAccessControlAllowOrigins `json:"accessControlAllowOrigins,omitempty" yaml:"accessControlAllowOrigins,omitempty"`

	// Object that contains an attribute `items` that contains a list of HTTP headers that CloudFront includes as values for the `Access-Control-Expose-Headers` HTTP response header.
	AccessControlExposeHeaders Cloudfront_ResponseHeadersPolicyCorsConfigAccessControlExposeHeaders `json:"accessControlExposeHeaders,omitempty" yaml:"accessControlExposeHeaders,omitempty"`
}
