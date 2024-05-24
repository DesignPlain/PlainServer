package types

type Cloudfront_DistributionOrderedCacheBehavior struct {
	// List of AWS account IDs (or `self`) that you want to allow to create signed URLs for private content. See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
	TrustedSigners []string `json:"trustedSigners,omitempty" yaml:"trustedSigners,omitempty"`

	// A config block that triggers a cloudfront function with specific actions (maximum 2).
	FunctionAssociations []Cloudfront_DistributionOrderedCacheBehaviorFunctionAssociation `json:"functionAssociations,omitempty" yaml:"functionAssociations,omitempty"`

	// Minimum amount of time that you want objects to stay in CloudFront caches before CloudFront queries your origin to see whether the object has been updated. Defaults to 0 seconds.
	MinTtl int `json:"minTtl,omitempty" yaml:"minTtl,omitempty"`

	// Pattern (for example, `images/-.jpg`) that specifies which requests you want this cache behavior to apply to.
	PathPattern string `json:"pathPattern,omitempty" yaml:"pathPattern,omitempty"`

	// List of key group IDs that CloudFront can use to validate signed URLs or signed cookies. See the [CloudFront User Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-trusted-signers.html) for more information about this feature.
	TrustedKeyGroups []string `json:"trustedKeyGroups,omitempty" yaml:"trustedKeyGroups,omitempty"`

	// Indicates whether you want to distribute media files in Microsoft Smooth Streaming format using the origin that is associated with this cache behavior.
	SmoothStreaming bool `json:"smoothStreaming,omitempty" yaml:"smoothStreaming,omitempty"`

	// Identifier for a response headers policy.
	ResponseHeadersPolicyId string `json:"responseHeadersPolicyId,omitempty" yaml:"responseHeadersPolicyId,omitempty"`

	// Value of ID for the origin that you want CloudFront to route requests to when a request matches the path pattern either for a cache behavior or for the default cache behavior.
	TargetOriginId string `json:"targetOriginId,omitempty" yaml:"targetOriginId,omitempty"`

	// Unique identifier of the cache policy that is attached to the cache behavior. If configuring the `default_cache_behavior` either `cache_policy_id` or `forwarded_values` must be set.
	CachePolicyId string `json:"cachePolicyId,omitempty" yaml:"cachePolicyId,omitempty"`

	// Maximum amount of time (in seconds) that an object is in a CloudFront cache before CloudFront forwards another request to your origin to determine whether the object has been updated. Only effective in the presence of `Cache-Control max-age`, `Cache-Control s-maxage`, and `Expires` headers.
	MaxTtl int `json:"maxTtl,omitempty" yaml:"maxTtl,omitempty"`

	// Unique identifier of the origin request policy that is attached to the behavior.
	OriginRequestPolicyId string `json:"originRequestPolicyId,omitempty" yaml:"originRequestPolicyId,omitempty"`

	// ARN of the real-time log configuration that is attached to this cache behavior.
	RealtimeLogConfigArn string `json:"realtimeLogConfigArn,omitempty" yaml:"realtimeLogConfigArn,omitempty"`

	// Field level encryption configuration ID.
	FieldLevelEncryptionId string `json:"fieldLevelEncryptionId,omitempty" yaml:"fieldLevelEncryptionId,omitempty"`

	// The forwarded values configuration that specifies how CloudFront handles query strings, cookies and headers (maximum one).
	ForwardedValues Cloudfront_DistributionOrderedCacheBehaviorForwardedValues `json:"forwardedValues,omitempty" yaml:"forwardedValues,omitempty"`

	// A config block that triggers a lambda function with specific actions (maximum 4).
	LambdaFunctionAssociations []Cloudfront_DistributionOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociations,omitempty" yaml:"lambdaFunctionAssociations,omitempty"`

	// Use this element to specify the protocol that users can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern. One of `allow-all`, `https-only`, or `redirect-to-https`.
	ViewerProtocolPolicy string `json:"viewerProtocolPolicy,omitempty" yaml:"viewerProtocolPolicy,omitempty"`

	// Controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin.
	AllowedMethods []string `json:"allowedMethods,omitempty" yaml:"allowedMethods,omitempty"`

	// Controls whether CloudFront caches the response to requests using the specified HTTP methods.
	CachedMethods []string `json:"cachedMethods,omitempty" yaml:"cachedMethods,omitempty"`

	// Whether you want CloudFront to automatically compress content for web requests that include `Accept-Encoding: gzip` in the request header (default: `false`).
	Compress bool `json:"compress,omitempty" yaml:"compress,omitempty"`

	// Default amount of time (in seconds) that an object is in a CloudFront cache before CloudFront forwards another request in the absence of an `Cache-Control max-age` or `Expires` header.
	DefaultTtl int `json:"defaultTtl,omitempty" yaml:"defaultTtl,omitempty"`
}
