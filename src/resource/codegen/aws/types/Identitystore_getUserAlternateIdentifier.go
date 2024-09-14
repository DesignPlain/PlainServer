package types

type Identitystore_getUserAlternateIdentifier struct {
	// Configuration block for filtering by the identifier issued by an external identity provider. Detailed below.
	ExternalId Identitystore_getUserAlternateIdentifierExternalId `json:"externalId,omitempty" yaml:"externalId,omitempty"`

	/*
	   An entity attribute that's unique to a specific entity. Detailed below.

	   > Exactly one of the above arguments must be provided.
	*/
	UniqueAttribute Identitystore_getUserAlternateIdentifierUniqueAttribute `json:"uniqueAttribute,omitempty" yaml:"uniqueAttribute,omitempty"`
}
