package types

type Identityplatform_ConfigMultiTenant struct {
	// Whether this project can have tenants or not.
	AllowTenants bool `json:"allowTenants,omitempty" yaml:"allowTenants,omitempty"`

	/*
	   The default cloud parent org or folder that the tenant project should be created under.
	   The parent resource name should be in the format of "/", such as "folders/123" or "organizations/456".
	   If the value is not set, the tenant will be created under the same organization or folder as the agent project.
	*/
	DefaultTenantLocation string `json:"defaultTenantLocation,omitempty" yaml:"defaultTenantLocation,omitempty"`
}
