package grafana

type WorkspaceSamlConfiguration struct {
	// The IDP Metadata URL. Note that either `idp_metadata_url` or `idp_metadata_xml` (but not both) must be specified.
	IdpMetadataUrl string `json:"idpMetadataUrl,omitempty" yaml:"idpMetadataUrl,omitempty"`

	// The login validity duration.
	LoginValidityDuration int `json:"loginValidityDuration,omitempty" yaml:"loginValidityDuration,omitempty"`

	// The org assertion.
	OrgAssertion string `json:"orgAssertion,omitempty" yaml:"orgAssertion,omitempty"`

	// The editor role values.
	EditorRoleValues []string `json:"editorRoleValues,omitempty" yaml:"editorRoleValues,omitempty"`

	// The groups assertion.
	GroupsAssertion string `json:"groupsAssertion,omitempty" yaml:"groupsAssertion,omitempty"`

	// The email assertion.
	EmailAssertion string `json:"emailAssertion,omitempty" yaml:"emailAssertion,omitempty"`

	// The IDP Metadata XML. Note that either `idp_metadata_url` or `idp_metadata_xml` (but not both) must be specified.
	IdpMetadataXml string `json:"idpMetadataXml,omitempty" yaml:"idpMetadataXml,omitempty"`

	// The login assertion.
	LoginAssertion string `json:"loginAssertion,omitempty" yaml:"loginAssertion,omitempty"`

	// The name assertion.
	NameAssertion string `json:"nameAssertion,omitempty" yaml:"nameAssertion,omitempty"`

	// The role assertion.
	RoleAssertion string `json:"roleAssertion,omitempty" yaml:"roleAssertion,omitempty"`

	/*
	   The workspace id.

	   The following arguments are optional:
	*/
	WorkspaceId string `json:"workspaceId,omitempty" yaml:"workspaceId,omitempty"`

	// The admin role values.
	AdminRoleValues []string `json:"adminRoleValues,omitempty" yaml:"adminRoleValues,omitempty"`

	// The allowed organizations.
	AllowedOrganizations []string `json:"allowedOrganizations,omitempty" yaml:"allowedOrganizations,omitempty"`
}
