package types

type Container_AzureClusterAzureServicesAuthentication struct {
	// The Azure Active Directory Application ID for Authentication configuration.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// The Azure Active Directory Tenant ID for Authentication configuration.
	TenantId string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`
}
