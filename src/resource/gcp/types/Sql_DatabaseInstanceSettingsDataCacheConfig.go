package types

type Sql_DatabaseInstanceSettingsDataCacheConfig struct {
	// Whether data cache is enabled for the instance. Defaults to `false`. Can be used with MYSQL and PostgreSQL only.
	DataCacheEnabled bool `json:"dataCacheEnabled,omitempty" yaml:"dataCacheEnabled,omitempty"`
}
