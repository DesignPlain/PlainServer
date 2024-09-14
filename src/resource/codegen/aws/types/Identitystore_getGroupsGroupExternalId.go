package types

type Identitystore_getGroupsGroupExternalId struct {
	// Identifier issued to this resource by an external identity provider.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Issuer for an external identifier.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
