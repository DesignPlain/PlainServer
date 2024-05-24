package types

type Opensearch_DomainAdvancedSecurityOptions struct {
	// Whether Anonymous auth is enabled. Enables fine-grained access control on an existing domain. Ignored unless `advanced_security_options` are enabled. _Can only be enabled on an existing domain._
	AnonymousAuthEnabled bool `json:"anonymousAuthEnabled,omitempty" yaml:"anonymousAuthEnabled,omitempty"`

	// Whether advanced security is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the internal user database is enabled. Default is `false`.
	InternalUserDatabaseEnabled bool `json:"internalUserDatabaseEnabled,omitempty" yaml:"internalUserDatabaseEnabled,omitempty"`

	// Configuration block for the main user. Detailed below.
	MasterUserOptions Opensearch_DomainAdvancedSecurityOptionsMasterUserOptions `json:"masterUserOptions,omitempty" yaml:"masterUserOptions,omitempty"`
}
