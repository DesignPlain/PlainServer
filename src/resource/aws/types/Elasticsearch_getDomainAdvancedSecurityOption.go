package types

type Elasticsearch_getDomainAdvancedSecurityOption struct {
	// Whether node to node encryption is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the internal user database is enabled.
	InternalUserDatabaseEnabled bool `json:"internalUserDatabaseEnabled,omitempty" yaml:"internalUserDatabaseEnabled,omitempty"`
}
