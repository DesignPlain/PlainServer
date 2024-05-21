package types

type Compute_getBackendBucketCdnPolicy struct {
	/*
	   Specifies the cache setting for all responses from this backend.
	   The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]
	*/
	CacheMode string `json:"cacheMode,omitempty" yaml:"cacheMode,omitempty"`

	// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
	NegativeCaching bool `json:"negativeCaching,omitempty" yaml:"negativeCaching,omitempty"`

	/*
	   Specifies the default TTL for cached content served by this origin for responses
	   that do not have an existing valid TTL (max-age or s-max-age).
	*/
	DefaultTtl int `json:"defaultTtl,omitempty" yaml:"defaultTtl,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	MaxTtl int `json:"maxTtl,omitempty" yaml:"maxTtl,omitempty"`

	/*
	   Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	   Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
	*/
	NegativeCachingPolicies []Compute_getBackendBucketCdnPolicyNegativeCachingPolicy `json:"negativeCachingPolicies,omitempty" yaml:"negativeCachingPolicies,omitempty"`

	// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
	RequestCoalescing bool `json:"requestCoalescing,omitempty" yaml:"requestCoalescing,omitempty"`

	// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
	ServeWhileStale int `json:"serveWhileStale,omitempty" yaml:"serveWhileStale,omitempty"`

	// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
	BypassCacheOnRequestHeaders []Compute_getBackendBucketCdnPolicyBypassCacheOnRequestHeader `json:"bypassCacheOnRequestHeaders,omitempty" yaml:"bypassCacheOnRequestHeaders,omitempty"`

	// The CacheKeyPolicy for this CdnPolicy.
	CacheKeyPolicies []Compute_getBackendBucketCdnPolicyCacheKeyPolicy `json:"cacheKeyPolicies,omitempty" yaml:"cacheKeyPolicies,omitempty"`

	// Specifies the maximum allowed TTL for cached content served by this origin.
	ClientTtl int `json:"clientTtl,omitempty" yaml:"clientTtl,omitempty"`

	/*
	   Maximum number of seconds the response to a signed URL request will
	   be considered fresh. After this time period,
	   the response will be revalidated before being served.
	   When serving responses to signed URL requests,
	   Cloud CDN will internally behave as though
	   all responses from this backend had a "Cache-Control: public,
	   max-age=[TTL]" header, regardless of any existing Cache-Control
	   header. The actual headers served in responses will not be altered.
	*/
	SignedUrlCacheMaxAgeSec int `json:"signedUrlCacheMaxAgeSec,omitempty" yaml:"signedUrlCacheMaxAgeSec,omitempty"`
}
