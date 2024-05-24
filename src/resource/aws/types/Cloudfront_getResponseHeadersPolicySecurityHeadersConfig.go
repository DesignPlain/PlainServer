package types

type Cloudfront_getResponseHeadersPolicySecurityHeadersConfig struct {
	// Value of the Referrer-Policy HTTP response header. Valid Values: `no-referrer` | `no-referrer-when-downgrade` | `origin` | `origin-when-cross-origin` | `same-origin` | `strict-origin` | `strict-origin-when-cross-origin` | `unsafe-url`
	ReferrerPolicies []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigReferrerPolicy `json:"referrerPolicies,omitempty" yaml:"referrerPolicies,omitempty"`

	// Settings that determine whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value. See Strict Transport Security for more information.
	StrictTransportSecurities []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity `json:"strictTransportSecurities,omitempty" yaml:"strictTransportSecurities,omitempty"`

	// Settings that determine whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value. See XSS Protection for more information.
	XssProtections []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigXssProtection `json:"xssProtections,omitempty" yaml:"xssProtections,omitempty"`

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header.
	ContentSecurityPolicies []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy `json:"contentSecurityPolicies,omitempty" yaml:"contentSecurityPolicies,omitempty"`

	// A setting that determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff. See Content Type Options for more information.
	ContentTypeOptions []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigContentTypeOption `json:"contentTypeOptions,omitempty" yaml:"contentTypeOptions,omitempty"`

	// Setting that determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value. See Frame Options for more information.
	FrameOptions []Cloudfront_getResponseHeadersPolicySecurityHeadersConfigFrameOption `json:"frameOptions,omitempty" yaml:"frameOptions,omitempty"`
}
