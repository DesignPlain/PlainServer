package types

type Opensearch_getDomainAdvancedSecurityOption struct {
	//
	AnonymousAuthEnabled bool `json:"anonymousAuthEnabled,omitempty" yaml:"anonymousAuthEnabled,omitempty"`

	// Enabled disabled toggle for off-peak update window
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the internal user database is enabled.
	InternalUserDatabaseEnabled bool `json:"internalUserDatabaseEnabled,omitempty" yaml:"internalUserDatabaseEnabled,omitempty"`
}
