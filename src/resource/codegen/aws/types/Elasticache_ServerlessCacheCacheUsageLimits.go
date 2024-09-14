package types

type Elasticache_ServerlessCacheCacheUsageLimits struct {
	// The maximum data storage limit in the cache, expressed in Gigabytes. See Data Storage config for more details.
	DataStorage Elasticache_ServerlessCacheCacheUsageLimitsDataStorage `json:"dataStorage,omitempty" yaml:"dataStorage,omitempty"`

	// The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second.See config block for more details.
	EcpuPerSeconds []Elasticache_ServerlessCacheCacheUsageLimitsEcpuPerSecond `json:"ecpuPerSeconds,omitempty" yaml:"ecpuPerSeconds,omitempty"`
}
