package compute

type OrganizationSecurityPolicy struct {
	/*
	   The type indicates the intended use of the security policy.
	   For organization security policies, the only supported type
	   is "FIREWALL".
	   Default value is `FIREWALL`.
	   Possible values are: `FIREWALL`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// A textual description for the organization security policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A textual name of the security policy.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	   Format: organizations/{organization_id} or folders/{folder_id}


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
