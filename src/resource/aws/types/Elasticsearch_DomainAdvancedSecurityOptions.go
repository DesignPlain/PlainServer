package types

type Elasticsearch_DomainAdvancedSecurityOptions struct {
	// Whether advanced security is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the internal user database is enabled. If not set, defaults to `false` by the AWS API.
	InternalUserDatabaseEnabled bool `json:"internalUserDatabaseEnabled,omitempty" yaml:"internalUserDatabaseEnabled,omitempty"`

	// Configuration block for the main user. Detailed below.
	MasterUserOptions Elasticsearch_DomainAdvancedSecurityOptionsMasterUserOptions `json:"masterUserOptions,omitempty" yaml:"masterUserOptions,omitempty"`
}
