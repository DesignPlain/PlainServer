package appsync

type ApiCache struct {
	// Cache instance type. Valid values are `SMALL`, `MEDIUM`, `LARGE`, `XLARGE`, `LARGE_2X`, `LARGE_4X`, `LARGE_8X`, `LARGE_12X`, `T2_SMALL`, `T2_MEDIUM`, `R4_LARGE`, `R4_XLARGE`, `R4_2XLARGE`, `R4_4XLARGE`, `R4_8XLARGE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Caching behavior. Valid values are `FULL_REQUEST_CACHING` and `PER_RESOLVER_CACHING`.
	ApiCachingBehavior string `json:"apiCachingBehavior,omitempty" yaml:"apiCachingBehavior,omitempty"`

	// GraphQL API ID.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// At-rest encryption flag for cache. You cannot update this setting after creation.
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" yaml:"atRestEncryptionEnabled,omitempty"`

	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled,omitempty" yaml:"transitEncryptionEnabled,omitempty"`

	// TTL in seconds for cache entries.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
