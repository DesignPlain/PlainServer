package container

type AzureClient struct {
	// The Azure Active Directory Application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of this resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The Azure Active Directory Tenant ID.



	   - - -
	*/
	TenantId string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`
}
