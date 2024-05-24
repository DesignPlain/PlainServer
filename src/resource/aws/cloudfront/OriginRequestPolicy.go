package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type OriginRequestPolicy struct {
	// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
	QueryStringsConfig types.Cloudfront_OriginRequestPolicyQueryStringsConfig `json:"queryStringsConfig,omitempty" yaml:"queryStringsConfig,omitempty"`

	// Comment to describe the origin request policy.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
	CookiesConfig types.Cloudfront_OriginRequestPolicyCookiesConfig `json:"cookiesConfig,omitempty" yaml:"cookiesConfig,omitempty"`

	// Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
	HeadersConfig types.Cloudfront_OriginRequestPolicyHeadersConfig `json:"headersConfig,omitempty" yaml:"headersConfig,omitempty"`

	// Unique name to identify the origin request policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
