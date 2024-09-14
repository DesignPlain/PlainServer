package types

type Identitystore_getGroupsGroup struct {
	// Identity Store ID associated with the Single Sign-On (SSO) Instance.
	IdentityStoreId string `json:"identityStoreId,omitempty" yaml:"identityStoreId,omitempty"`

	// Description of the specified group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Group's display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// List of identifiers issued to this resource by an external identity provider.
	ExternalIds []Identitystore_getGroupsGroupExternalId `json:"externalIds,omitempty" yaml:"externalIds,omitempty"`

	// Identifier of the group in the Identity Store.
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`
}
