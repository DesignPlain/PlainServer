package identitystore

type Group struct {
	// A string containing the description of the group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A string containing the name of the group. This value is commonly displayed when the group is referenced.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The globally unique identifier for the identity store.

	   The following arguments are optional:
	*/
	IdentityStoreId string `json:"identityStoreId,omitempty" yaml:"identityStoreId,omitempty"`
}
