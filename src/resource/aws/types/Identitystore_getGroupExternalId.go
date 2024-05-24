package types

type Identitystore_getGroupExternalId struct {
	// The identifier issued to this resource by an external identity provider.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The issuer for an external identifier.
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty"`
}
