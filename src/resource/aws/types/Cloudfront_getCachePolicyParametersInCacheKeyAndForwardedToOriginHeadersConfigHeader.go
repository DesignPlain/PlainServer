package types

type Cloudfront_getCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeader struct {
	// List of item names (`cookies`, `headers`, or `query_strings`).
	Items []string `json:"items,omitempty" yaml:"items,omitempty"`
}
