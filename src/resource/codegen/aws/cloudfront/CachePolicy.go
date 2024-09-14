package cloudfront

import types "libds/aws/types"

type CachePolicy struct {
	// Amount of time, in seconds, that objects are allowed to remain in the CloudFront cache before CloudFront sends a new request to the origin server to check if the object has been updated.
	DefaultTtl int `json:"defaultTtl,omitempty" yaml:"defaultTtl,omitempty"`

	// Maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl int `json:"maxTtl,omitempty" yaml:"maxTtl,omitempty"`

	// Minimum amount of time, in seconds, that objects should remain in the CloudFront cache before a new request is sent to the origin to check for updates.
	MinTtl int `json:"minTtl,omitempty" yaml:"minTtl,omitempty"`

	// Unique name used to identify the cache policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration for including HTTP headers, cookies, and URL query strings in the cache key. For more information, refer to the Parameters In Cache Key And Forwarded To Origin section.
	ParametersInCacheKeyAndForwardedToOrigin types.Cloudfront_CachePolicyParametersInCacheKeyAndForwardedToOrigin `json:"parametersInCacheKeyAndForwardedToOrigin,omitempty" yaml:"parametersInCacheKeyAndForwardedToOrigin,omitempty"`

	// Description for the cache policy.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`
}
