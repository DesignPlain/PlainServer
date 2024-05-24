package types

type Cloudfront_DistributionCustomErrorResponse struct {
	// 4xx or 5xx HTTP status code that you want to customize.
	ErrorCode int `json:"errorCode,omitempty" yaml:"errorCode,omitempty"`

	// HTTP status code that you want CloudFront to return with the custom error page to the viewer.
	ResponseCode int `json:"responseCode,omitempty" yaml:"responseCode,omitempty"`

	// Path of the custom error page (for example, `/custom_404.html`).
	ResponsePagePath string `json:"responsePagePath,omitempty" yaml:"responsePagePath,omitempty"`

	// Minimum amount of time you want HTTP error codes to stay in CloudFront caches before CloudFront queries your origin to see whether the object has been updated.
	ErrorCachingMinTtl int `json:"errorCachingMinTtl,omitempty" yaml:"errorCachingMinTtl,omitempty"`
}
