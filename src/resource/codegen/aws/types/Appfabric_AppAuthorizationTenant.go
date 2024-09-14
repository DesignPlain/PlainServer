package types

type Appfabric_AppAuthorizationTenant struct {
	// The display name of the tenant.
	TenantDisplayName string `json:"tenantDisplayName,omitempty" yaml:"tenantDisplayName,omitempty"`

	// The ID of the application tenant.
	TenantIdentifier string `json:"tenantIdentifier,omitempty" yaml:"tenantIdentifier,omitempty"`
}
