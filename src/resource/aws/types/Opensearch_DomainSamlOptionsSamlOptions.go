package types

type Opensearch_DomainSamlOptionsSamlOptions struct {
	// Whether SAML authentication is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Information from your identity provider.
	Idp Opensearch_DomainSamlOptionsSamlOptionsIdp `json:"idp,omitempty" yaml:"idp,omitempty"`

	// This backend role from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterBackendRole string `json:"masterBackendRole,omitempty" yaml:"masterBackendRole,omitempty"`

	// This username from the SAML IdP receives full permissions to the cluster, equivalent to a new master user.
	MasterUserName string `json:"masterUserName,omitempty" yaml:"masterUserName,omitempty"`

	// Element of the SAML assertion to use for backend roles. Default is roles.
	RolesKey string `json:"rolesKey,omitempty" yaml:"rolesKey,omitempty"`

	// Duration of a session in minutes after a user logs in. Default is 60. Maximum value is 1,440.
	SessionTimeoutMinutes int `json:"sessionTimeoutMinutes,omitempty" yaml:"sessionTimeoutMinutes,omitempty"`

	// Element of the SAML assertion to use for username. Default is NameID.
	SubjectKey string `json:"subjectKey,omitempty" yaml:"subjectKey,omitempty"`
}
