package types

type Elasticache_ServerlessCacheCacheUsageLimitsDataStorage struct {
	// The unit that the storage is measured in, in GB.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// The upper limit for data storage the cache is set to use. Set as Integer.
	Maximum int `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}
