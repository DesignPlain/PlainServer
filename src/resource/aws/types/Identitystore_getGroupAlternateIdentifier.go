package types

type Identitystore_getGroupAlternateIdentifier struct {
	// Configuration block for filtering by the identifier issued by an external identity provider. Detailed below.
	ExternalId Identitystore_getGroupAlternateIdentifierExternalId `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	/*
	   An entity attribute that's unique to a specific entity. Detailed below.

	   > Exactly one of the above arguments must be provided.
	*/
	UniqueAttribute Identitystore_getGroupAlternateIdentifierUniqueAttribute `json:"uniqueAttribute,omitempty" yaml:"uniqueAttribute,omitempty"`
}
