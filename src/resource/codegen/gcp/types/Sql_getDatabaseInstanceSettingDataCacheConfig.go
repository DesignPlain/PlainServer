package types

type Sql_getDatabaseInstanceSettingDataCacheConfig struct {
	// Whether data cache is enabled for the instance.
	DataCacheEnabled bool `json:"dataCacheEnabled,omitempty" yaml:"dataCacheEnabled,omitempty"`
}
