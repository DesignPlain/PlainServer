package types

type Finspace_KxClusterDatabaseCacheConfiguration struct {
	// Type of disk cache.
	CacheType string `json:"cacheType,omitempty" yaml:"cacheType,omitempty"`

	// Paths within the database to cache.
	DbPaths []string `json:"dbPaths,omitempty" yaml:"dbPaths,omitempty"`
}
