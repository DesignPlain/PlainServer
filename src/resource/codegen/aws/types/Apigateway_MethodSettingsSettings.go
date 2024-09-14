package types

type Apigateway_MethodSettingsSettings struct {
	// Whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled bool `json:"metricsEnabled,omitempty" yaml:"metricsEnabled,omitempty"`

	// Whether authorization is required for a cache invalidation request.
	RequireAuthorizationForCacheControl bool `json:"requireAuthorizationForCacheControl,omitempty" yaml:"requireAuthorizationForCacheControl,omitempty"`

	// Throttling burst limit. Default: `-1` (throttling disabled).
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" yaml:"throttlingBurstLimit,omitempty"`

	// Throttling rate limit. Default: `-1` (throttling disabled).
	ThrottlingRateLimit float64 `json:"throttlingRateLimit,omitempty" yaml:"throttlingRateLimit,omitempty"`

	// How to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.
	UnauthorizedCacheControlHeaderStrategy string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" yaml:"unauthorizedCacheControlHeaderStrategy,omitempty"`

	// Whether the cached responses are encrypted.
	CacheDataEncrypted bool `json:"cacheDataEncrypted,omitempty" yaml:"cacheDataEncrypted,omitempty"`

	// Whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled bool `json:"cachingEnabled,omitempty" yaml:"cachingEnabled,omitempty"`

	// Whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" yaml:"dataTraceEnabled,omitempty"`

	// Time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
	CacheTtlInSeconds int `json:"cacheTtlInSeconds,omitempty" yaml:"cacheTtlInSeconds,omitempty"`

	// Logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
	LoggingLevel string `json:"loggingLevel,omitempty" yaml:"loggingLevel,omitempty"`
}
