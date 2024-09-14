package types

type Elasticache_ServerlessCacheCacheUsageLimitsDataStorage struct {
	//
	Minimum int `json:"minimum,omitempty" yaml:"minimum,omitempty"`

	// The unit that the storage is measured in, in GB.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	//
	Maximum int `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}
