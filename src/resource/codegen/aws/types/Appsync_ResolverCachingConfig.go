package types

type Appsync_ResolverCachingConfig struct {
	// The caching keys for a resolver that has caching activated. Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	CachingKeys []string `json:"cachingKeys,omitempty" yaml:"cachingKeys,omitempty"`

	// The TTL in seconds for a resolver that has caching activated. Valid values are between `1` and `3600` seconds.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
