package types

type Cloudfront_ResponseHeadersPolicySecurityHeadersConfig struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header. See Content Security Policy for more information.
	ContentSecurityPolicy Cloudfront_ResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy `json:"contentSecurityPolicy,omitempty" yaml:"contentSecurityPolicy,omitempty"`

	// Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response header with its value set to `nosniff`. See Content Type Options for more information.
	ContentTypeOptions Cloudfront_ResponseHeadersPolicySecurityHeadersConfigContentTypeOptions `json:"contentTypeOptions,omitempty" yaml:"contentTypeOptions,omitempty"`

	// Determines whether CloudFront includes the `X-Frame-Options` HTTP response header and the header’s value. See Frame Options for more information.
	FrameOptions Cloudfront_ResponseHeadersPolicySecurityHeadersConfigFrameOptions `json:"frameOptions,omitempty" yaml:"frameOptions,omitempty"`

	// Determines whether CloudFront includes the `Referrer-Policy` HTTP response header and the header’s value. See Referrer Policy for more information.
	ReferrerPolicy Cloudfront_ResponseHeadersPolicySecurityHeadersConfigReferrerPolicy `json:"referrerPolicy,omitempty" yaml:"referrerPolicy,omitempty"`

	// Determines whether CloudFront includes the `Strict-Transport-Security` HTTP response header and the header’s value. See Strict Transport Security for more information.
	StrictTransportSecurity Cloudfront_ResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity `json:"strictTransportSecurity,omitempty" yaml:"strictTransportSecurity,omitempty"`

	// Determine whether CloudFront includes the `X-XSS-Protection` HTTP response header and the header’s value. See XSS Protection for more information.
	XssProtection Cloudfront_ResponseHeadersPolicySecurityHeadersConfigXssProtection `json:"xssProtection,omitempty" yaml:"xssProtection,omitempty"`
}
